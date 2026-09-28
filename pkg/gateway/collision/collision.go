// Copyright 2019-2026 The Liqo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package collision applies a routing rule on the gateway pod that diverts
// fabric (geneve) traffic directed to the gateway pod IP through the tunnel.
// This is required when the local and remote clusters share the same pod CIDR,
// because in that case the gateway pod IP belongs to the remote pod CIDR and
// would otherwise be delivered locally. The rule is applied unconditionally:
// when CIDRs do not overlap it is still harmless because the destination is
// the gateway pod IP itself.
package collision

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"

	networkingv1beta1 "github.com/liqotech/liqo/apis/networking/v1beta1"
	"github.com/liqotech/liqo/pkg/gateway"
	"github.com/liqotech/liqo/pkg/gateway/tunnel"
	"github.com/liqotech/liqo/pkg/route"
)

const (
	// gwExtMark is the fwmark value used to tag traffic arriving on Geneve
	// interfaces. It must match the mark set by the FirewallConfiguration
	// created by the controller-manager.
	gwExtMark = 0xFF00

	// collisionRulePriority is the priority of the rule diverting fabric traffic
	// directed to the gateway pod IP through the tunnel. It must be lower (more
	// preferred) than the priority of the local routing rule (10, see
	// kernel.DeprioritizeLocalRule), so that this rule is evaluated first.
	collisionRulePriority = 0

	// podIPEnvVar is the name of the environment variable injected via downward
	// API containing the pod IP.
	podIPEnvVar = "POD_IP"

	// routeRetryInterval is the interval between retries when the collision
	// route cannot be installed because the gateway is not reachable yet.
	routeRetryInterval = 2 * time.Second
)

// ApplyGatewayCollisionRule applies the collision rule on the gateway pod. The
// rule is created immediately, even if the routing table does not exist yet.
// The host route in the table is retried until the remote tunnel gateway
// becomes reachable.
func ApplyGatewayCollisionRule(ctx context.Context, opts *gateway.Options) error {
	gwPodIP := os.Getenv(podIPEnvVar)
	if gwPodIP == "" {
		return fmt.Errorf("environment variable %s is not set", podIPEnvVar)
	}
	klog.Infof("Applying gateway collision rule for pod IP %s", gwPodIP)

	remoteInterfaceIP, err := tunnel.GetRemoteInterfaceIP(opts.Mode)
	if err != nil {
		return err
	}
	klog.V(4).Infof("Remote tunnel interface IP: %s", remoteInterfaceIP)

	tableID, err := route.GetTableID(opts.Name)
	if err != nil {
		return fmt.Errorf("getting the table ID: %w", err)
	}
	klog.V(4).Infof("Routing table ID: %d", tableID)

	gwPodIPCIDR := networkingv1beta1.CIDR(gwPodIP + "/32")
	mark := gwExtMark
	priority := collisionRulePriority

	desiredRule := networkingv1beta1.Rule{
		FwMark:   &mark,
		Dst:      &gwPodIPCIDR,
		Priority: &priority,
		Routes: []networkingv1beta1.Route{
			{
				Dst: &gwPodIPCIDR,
				Gw:  ptr.To(networkingv1beta1.IP(remoteInterfaceIP)),
			},
		},
	}

	existingRules, err := route.GetRulesByTableID(tableID)
	if err != nil {
		return fmt.Errorf("listing existing rules: %w", err)
	}

	if err := route.EnsureRulePresence(&desiredRule, tableID, existingRules); err != nil {
		return fmt.Errorf("ensuring rule presence: %w", err)
	}

	klog.Infof("Applied gateway collision rule (pod IP %s, table %d, priority %d)", gwPodIP, tableID, priority)

	// The host route may fail if the remote tunnel gateway is not reachable
	// yet (e.g., the tunnel interface is still being configured). Retry until
	// it succeeds or the context is cancelled.
	if err := ensureRouteWithRetry(ctx, desiredRule.Routes, tableID); err != nil {
		return fmt.Errorf("ensuring routes presence: %w", err)
	}

	klog.Infof("Applied gateway collision route (pod IP %s, gateway %s, table %d)", gwPodIP, remoteInterfaceIP, tableID)
	return nil
}

// ensureRouteWithRetry keeps trying to install the collision route until it
// succeeds or the context is cancelled.
func ensureRouteWithRetry(ctx context.Context, routes []networkingv1beta1.Route, tableID uint32) error {
	attempts := 0
	for {
		attempts++
		existingRoutes, err := route.GetRoutesByTableID(tableID)
		if err != nil {
			return fmt.Errorf("listing existing routes: %w", err)
		}

		err = route.EnsureRoutesPresence(routes, tableID, existingRoutes)
		if err == nil {
			klog.V(4).Infof("Collision route installed after %d attempt(s)", attempts)
			return nil
		}
		if !errors.Is(err, route.ErrNetworkUnreachable) {
			return err
		}

		klog.V(3).Infof("Collision route not reachable yet (attempt %d), retrying in %v: %v", attempts, routeRetryInterval, err)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(routeRetryInterval):
		}
	}
}
