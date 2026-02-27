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

package route

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

// statusMapGetter is a function that returns a pointer to the RouteConfigurations status map.
type statusMapGetter func(obj client.Object) *map[string][]metav1.Condition

// buildRouteCondition creates a condition for the RouteConfiguration status.
func buildRouteCondition(rtcfg *networkingv1beta1.RouteConfiguration, applyErr error) metav1.Condition {
	conditionStatus := metav1.ConditionTrue
	if applyErr != nil {
		conditionStatus = metav1.ConditionFalse
	}
	return metav1.Condition{
		Type:               string(networkingv1beta1.RouteConfigurationStatusConditionTypeApplied),
		Status:             conditionStatus,
		LastTransitionTime: metav1.Now(),
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

// updateRouteConfigurationStatus is a generic function that updates RouteConfiguration status for any resource.
func updateRouteConfigurationStatus(
	ctx context.Context,
	cl client.Client,
	objKey client.ObjectKey,
	newObjFunc func() client.Object,
	gvk schema.GroupVersionKind,
	resourceKind string,
	fieldOwnerPrefix string,
	getStatusMap statusMapGetter,
	rtcfg *networkingv1beta1.RouteConfiguration,
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
	condition := buildRouteCondition(rtcfg, applyErr)

	// Create a patch object for Server-Side Apply
	patch := newObjFunc()
	patch.SetName(obj.GetName())
	patch.SetNamespace(obj.GetNamespace())
	patch.GetObjectKind().SetGroupVersionKind(gvk)

	// Copy existing status map and update the condition for this RouteConfiguration
	key := fmt.Sprintf("%s/%s", rtcfg.Namespace, rtcfg.Name)
	existingStatusMap := *getStatusMap(obj)

	// Create a new map with all existing entries to avoid overwriting other RouteConfigurations
	statusMap := getStatusMap(patch)
	*statusMap = make(map[string][]metav1.Condition, len(existingStatusMap))
	for k, v := range existingStatusMap {
		(*statusMap)[k] = v
	}

	// Update or add the condition for the current RouteConfiguration
	(*statusMap)[key] = updateOrAddCondition(existingStatusMap[key], condition)

	// Apply the status patch using Server-Side Apply
	patchOpts := []client.SubResourcePatchOption{
		client.FieldOwner(fieldOwnerPrefix),
		client.ForceOwnership,
	}
	if err := cl.Status().Patch(ctx, patch, client.Apply, patchOpts...); err != nil {
		return fmt.Errorf("unable to patch %s %s status: %w", resourceKind, objKey, err)
	}

	klog.V(4).Infof("Updated %s %s status for RouteConfiguration %s (field owner: %s)",
		resourceKind, objKey, rtcfg.Name, fieldOwnerPrefix)
	return nil
}

// NewInternalNodeStatusUpdateFunc creates a StatusUpdateFunc that updates InternalNode status.
func NewInternalNodeStatusUpdateFunc(cl client.Client, nodeName string) StatusUpdateFunc {
	return func(ctx context.Context, rtcfg *networkingv1beta1.RouteConfiguration, applyErr error) error {
		return updateRouteConfigurationStatus(
			ctx, cl,
			client.ObjectKey{Name: nodeName},
			func() client.Object { return &networkingv1beta1.InternalNode{} },
			networkingv1beta1.InternalNodeGroupVersionResource.GroupVersion().WithKind(networkingv1beta1.InternalNodeKind),
			"InternalNode",
			"liqo-fabric-route-controller",
			func(obj client.Object) *map[string][]metav1.Condition {
				return &obj.(*networkingv1beta1.InternalNode).Status.RouteConfigurations
			},
			rtcfg, applyErr,
		)
	}
}

// NewGatewayClientStatusUpdateFunc creates a StatusUpdateFunc that updates GatewayClient status.
func NewGatewayClientStatusUpdateFunc(cl client.Client, gatewayName, gatewayNamespace string) StatusUpdateFunc {
	return func(ctx context.Context, rtcfg *networkingv1beta1.RouteConfiguration, applyErr error) error {
		return updateRouteConfigurationStatus(
			ctx, cl,
			client.ObjectKey{Name: gatewayName, Namespace: gatewayNamespace},
			func() client.Object { return &networkingv1beta1.GatewayClient{} },
			networkingv1beta1.GatewayClientGroupVersionResource.GroupVersion().WithKind(networkingv1beta1.GatewayClientKind),
			"GatewayClient",
			"liqo-gateway-route-controller",
			func(obj client.Object) *map[string][]metav1.Condition {
				return &obj.(*networkingv1beta1.GatewayClient).Status.RouteConfigurations
			},
			rtcfg, applyErr,
		)
	}
}

// NewGatewayServerStatusUpdateFunc creates a StatusUpdateFunc that updates GatewayServer status.
func NewGatewayServerStatusUpdateFunc(cl client.Client, gatewayName, gatewayNamespace string) StatusUpdateFunc {
	return func(ctx context.Context, rtcfg *networkingv1beta1.RouteConfiguration, applyErr error) error {
		return updateRouteConfigurationStatus(
			ctx, cl,
			client.ObjectKey{Name: gatewayName, Namespace: gatewayNamespace},
			func() client.Object { return &networkingv1beta1.GatewayServer{} },
			networkingv1beta1.GatewayServerGroupVersionResource.GroupVersion().WithKind(networkingv1beta1.GatewayServerKind),
			"GatewayServer",
			"liqo-gateway-route-controller",
			func(obj client.Object) *map[string][]metav1.Condition {
				return &obj.(*networkingv1beta1.GatewayServer).Status.RouteConfigurations
			},
			rtcfg, applyErr,
		)
	}
}
