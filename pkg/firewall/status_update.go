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

package firewall

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	networkingv1beta1 "github.com/liqotech/liqo/apis/networking/v1beta1"
)

// statusMapGetter is a function that returns a pointer to the FirewallConfigurations status map.
type statusMapGetter func(obj client.Object) *map[string][]metav1.Condition

// buildFirewallCondition creates a condition for the FirewallConfiguration status.
func buildFirewallCondition(fwcfg *networkingv1beta1.FirewallConfiguration, applyErr error) metav1.Condition {
	conditionStatus := metav1.ConditionTrue
	if applyErr != nil {
		conditionStatus = metav1.ConditionFalse
	}
	return metav1.Condition{
		Type:               string(networkingv1beta1.FirewallConfigurationStatusConditionTypeApplied),
		Status:             conditionStatus,
		LastTransitionTime: metav1.Now(),
		Reason:             "",
		Message:            "",
	}
}

// updateOrAddCondition updates an existing condition or appends a new one to the list.
func updateOrAddCondition(conditions []metav1.Condition, newCondition metav1.Condition) []metav1.Condition {
	for i := range conditions {
		if conditions[i].Type == newCondition.Type {
			conditions[i] = newCondition
			return conditions
		}
	}
	return append(conditions, newCondition)
}

// updateFirewallConfigurationStatus is a generic function that updates FirewallConfiguration status for any resource.
func updateFirewallConfigurationStatus(
	ctx context.Context,
	cl client.Client,
	objKey client.ObjectKey,
	newObjFunc func() client.Object,
	gvk schema.GroupVersionKind,
	resourceKind string,
	fieldOwnerPrefix string,
	getStatusMap statusMapGetter,
	fwcfg *networkingv1beta1.FirewallConfiguration,
	applyErr error,
) error {
	// Get the resource
	obj := newObjFunc()
	if err := cl.Get(ctx, objKey, obj); err != nil {
		if apierrors.IsNotFound(err) {
			klog.V(4).Infof("%s %s not found, skipping status update", resourceKind, objKey)
			return nil
		}
		return fmt.Errorf("unable to get %s %s: %w", resourceKind, objKey, err)
	}

	// Create the condition
	condition := buildFirewallCondition(fwcfg, applyErr)

	// Create a patch object for Server-Side Apply
	patch := newObjFunc()
	patch.SetName(obj.GetName())
	patch.SetNamespace(obj.GetNamespace())
	patch.GetObjectKind().SetGroupVersionKind(gvk)

	// Copy existing status map and update the condition for this FirewallConfiguration
	key := fmt.Sprintf("%s/%s", fwcfg.Namespace, fwcfg.Name)
	existingStatusMap := *getStatusMap(obj)

	// Create a new map with all existing entries to avoid overwriting other FirewallConfigurations
	statusMap := getStatusMap(patch)
	*statusMap = make(map[string][]metav1.Condition, len(existingStatusMap))
	for k, v := range existingStatusMap {
		(*statusMap)[k] = v
	}

	// Update or add the condition for the current FirewallConfiguration
	(*statusMap)[key] = updateOrAddCondition(existingStatusMap[key], condition)

	// Apply the status patch using Server-Side Apply
	patchOpts := []client.SubResourcePatchOption{
		client.FieldOwner(fieldOwnerPrefix),
		client.ForceOwnership,
	}
	if err := cl.Status().Patch(ctx, patch, client.Apply, patchOpts...); err != nil {
		return fmt.Errorf("unable to patch %s %s status: %w", resourceKind, objKey, err)
	}

	klog.V(4).Infof("Updated %s %s status for FirewallConfiguration %s (field owner: %s)",
		resourceKind, objKey, fwcfg.Name, fieldOwnerPrefix)
	return nil
}

// NewInternalNodeStatusUpdateFunc creates a StatusUpdateFunc that updates InternalNode status.
func NewInternalNodeStatusUpdateFunc(cl client.Client, nodeName string) StatusUpdateFunc {
	return func(ctx context.Context, fwcfg *networkingv1beta1.FirewallConfiguration, applyErr error) error {
		return updateFirewallConfigurationStatus(
			ctx, cl,
			client.ObjectKey{Name: nodeName},
			func() client.Object { return &networkingv1beta1.InternalNode{} },
			networkingv1beta1.InternalNodeGroupVersionResource.GroupVersion().WithKind(networkingv1beta1.InternalNodeKind),
			"InternalNode",
			"liqo-fabric-firewall-controller",
			func(obj client.Object) *map[string][]metav1.Condition {
				return &obj.(*networkingv1beta1.InternalNode).Status.FirewallConfigurations
			},
			fwcfg, applyErr,
		)
	}
}

// NewGatewayClientStatusUpdateFunc creates a StatusUpdateFunc that updates GatewayClient status.
func NewGatewayClientStatusUpdateFunc(cl client.Client, gatewayName, gatewayNamespace string) StatusUpdateFunc {
	return func(ctx context.Context, fwcfg *networkingv1beta1.FirewallConfiguration, applyErr error) error {
		return updateFirewallConfigurationStatus(
			ctx, cl,
			client.ObjectKey{Name: gatewayName, Namespace: gatewayNamespace},
			func() client.Object { return &networkingv1beta1.GatewayClient{} },
			networkingv1beta1.GatewayClientGroupVersionResource.GroupVersion().WithKind(networkingv1beta1.GatewayClientKind),
			"GatewayClient",
			"liqo-gateway-firewall-controller",
			func(obj client.Object) *map[string][]metav1.Condition {
				return &obj.(*networkingv1beta1.GatewayClient).Status.FirewallConfigurations
			},
			fwcfg, applyErr,
		)
	}
}

// NewGatewayServerStatusUpdateFunc creates a StatusUpdateFunc that updates GatewayServer status.
func NewGatewayServerStatusUpdateFunc(cl client.Client, gatewayName, gatewayNamespace string) StatusUpdateFunc {
	return func(ctx context.Context, fwcfg *networkingv1beta1.FirewallConfiguration, applyErr error) error {
		return updateFirewallConfigurationStatus(
			ctx, cl,
			client.ObjectKey{Name: gatewayName, Namespace: gatewayNamespace},
			func() client.Object { return &networkingv1beta1.GatewayServer{} },
			networkingv1beta1.GatewayServerGroupVersionResource.GroupVersion().WithKind(networkingv1beta1.GatewayServerKind),
			"GatewayServer",
			"liqo-gateway-firewall-controller",
			func(obj client.Object) *map[string][]metav1.Condition {
				return &obj.(*networkingv1beta1.GatewayServer).Status.FirewallConfigurations
			},
			fwcfg, applyErr,
		)
	}
}
