// Copyright 2019-2022 The Liqo Authors
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

// Package getters provides functions to get k8s resources and liqo custom resources.
package getters

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"

	discoveryv1alpha "github.com/liqotech/liqo/apis/discovery/v1alpha1"
	netv1alpha1 "github.com/liqotech/liqo/apis/net/v1alpha1"
)

// GetIPAMStorageByLabel it returns a IPAMStorage instance that matches the given label selector.
func GetIPAMStorageByLabel(ctx context.Context, cl client.Client, ns string, lSelector labels.Selector) (*netv1alpha1.IpamStorage, error) {
	list := new(netv1alpha1.IpamStorageList)
	if err := cl.List(ctx, list, &client.ListOptions{LabelSelector: lSelector}, client.InNamespace(ns)); err != nil {
		return nil, err
	}

	switch len(list.Items) {
	case 0:
		return nil, kerrors.NewNotFound(netv1alpha1.IpamGroupResource, netv1alpha1.ResourceIpamStorages)
	case 1:
		return &list.Items[0], nil
	default:
		return nil, fmt.Errorf("multiple resources of type {%s} found for label selector {%s} in namespace {%s},"+
			" when only one was expected", netv1alpha1.IpamGroupResource.String(), lSelector.String(), ns)
	}
}

// GetNetworkConfigByLabel it returns a NetworkConfig instance that matches the given label selector.
func GetNetworkConfigByLabel(ctx context.Context, cl client.Client, ns string, lSelector labels.Selector) (*netv1alpha1.NetworkConfig, error) {
	list, err := GetNetworkConfigsByLabel(ctx, cl, ns, lSelector)
	if err != nil {
		return nil, err
	}

	switch len(list.Items) {
	case 0:
		return nil, kerrors.NewNotFound(netv1alpha1.NetworkConfigGroupResource, netv1alpha1.ResourceNetworkConfigs)
	case 1:
		return &list.Items[0], nil
	default:
		return nil, fmt.Errorf("multiple resources of type {%s} found for label selector {%s} in namespace {%s},"+
			" when only one was expected", netv1alpha1.NetworkConfigGroupResource.String(), lSelector.String(), ns)
	}
}

// GetNetworkConfigsByLabel it returns a NetworkConfig list that matches the given label selector.
func GetNetworkConfigsByLabel(ctx context.Context, cl client.Client, ns string, lSelector labels.Selector) (*netv1alpha1.NetworkConfigList, error) {
	list := new(netv1alpha1.NetworkConfigList)
	if err := cl.List(ctx, list, &client.ListOptions{LabelSelector: lSelector}, client.InNamespace(ns)); err != nil {
		return nil, err
	}
	return list, nil
}

// GetTunnelEndpointByLabel it returns a TunnelEndpoint instance that matches the given label selector.
func GetTunnelEndpointByLabel(ctx context.Context, cl client.Client, ns string, lSelector labels.Selector) (*netv1alpha1.TunnelEndpoint, error) {
	list, err := GetTunnelEndpointsByLabel(ctx, cl, ns, lSelector)
	if err != nil {
		return nil, err
	}
	switch len(list.Items) {
	case 0:
		return nil, kerrors.NewNotFound(netv1alpha1.TunnelEndpointGroupResource, netv1alpha1.ResourceTunnelEndpoints)
	case 1:
		return &list.Items[0], nil
	default:
		return nil, fmt.Errorf("multiple resources of type {%s} found for label selector {%s} in namespace {%s},"+
			" when only one was expected", netv1alpha1.TunnelEndpointGroupResource.String(), lSelector.String(), ns)
	}
}

// GetTunnelEndpointsByLabel it returns a TunnelEndpoint list that matches the given label selector.
func GetTunnelEndpointsByLabel(ctx context.Context, cl client.Client, ns string, lSelector labels.Selector) (*netv1alpha1.TunnelEndpointList, error) {
	list := new(netv1alpha1.TunnelEndpointList)
	if err := cl.List(ctx, list, &client.ListOptions{LabelSelector: lSelector}, client.InNamespace(ns)); err != nil {
		return nil, err
	}
	return list, nil
}

// GetForeignClustersByLabel it returns a ForeignClusters list that matches the given label selector.
func GetForeignClustersByLabel(ctx context.Context, cl client.Client, lSelector labels.Selector) (*discoveryv1alpha.ForeignClusterList, error) {
	list := new(discoveryv1alpha.ForeignClusterList)
	if err := cl.List(ctx, list, &client.ListOptions{LabelSelector: lSelector}); err != nil {
		return nil, err
	}
	if len(list.Items) == 0 {
		return nil, kerrors.NewNotFound(discoveryv1alpha.ForeignClusterGroupResource, discoveryv1alpha.ForeignClusterResource)
	}
	return list, nil
}

// GetServiceByLabel it returns a service instance that matches the given label selector.
func GetServiceByLabel(ctx context.Context, cl client.Client, ns string, lSelector labels.Selector) (*corev1.Service, error) {
	list := new(corev1.ServiceList)
	if err := cl.List(ctx, list, &client.ListOptions{LabelSelector: lSelector}, client.InNamespace(ns)); err != nil {
		return nil, err
	}
	svcRN := string(corev1.ResourceServices)
	svcGR := corev1.Resource(svcRN)

	switch len(list.Items) {
	case 0:
		return nil, kerrors.NewNotFound(svcGR, svcRN)
	case 1:
		return &list.Items[0], nil
	default:
		return nil, fmt.Errorf("multiple resources of type {%s} found for label selector {%s} in namespace {%s},"+
			" when only one was expected", svcGR.String(), lSelector.String(), ns)
	}
}

// GetSecretByLabel it returns a secret instance that matches the given label selector.
func GetSecretByLabel(ctx context.Context, cl client.Client, ns string, lSelector labels.Selector) (*corev1.Secret, error) {
	list := new(corev1.SecretList)
	if err := cl.List(ctx, list, &client.ListOptions{LabelSelector: lSelector}, client.InNamespace(ns)); err != nil {
		return nil, err
	}
	scrRN := string(corev1.ResourceSecrets)
	scrGR := corev1.Resource(scrRN)

	switch len(list.Items) {
	case 0:
		return nil, kerrors.NewNotFound(scrGR, scrRN)
	case 1:
		return &list.Items[0], nil
	default:
		return nil, fmt.Errorf("multiple resources of type {%s} found for label selector {%s} in namespace {%s},"+
			" when only one was expected", scrGR.String(), lSelector.String(), ns)
	}
}

// GetPodByLabel it returns a pod instance that matches the given label and field selector.
func GetPodByLabel(ctx context.Context, cl client.Client, ns string, lSelector labels.Selector, fSelector fields.Selector) (*corev1.Pod, error) {
	list := new(corev1.PodList)
	if err := cl.List(ctx, list, &client.ListOptions{
		LabelSelector: lSelector,
		FieldSelector: fSelector,
	}, client.InNamespace(ns)); err != nil {
		return nil, err
	}
	podRN := string(corev1.ResourcePods)
	podGR := corev1.Resource(podRN)

	switch len(list.Items) {
	case 0:
		return nil, kerrors.NewNotFound(podGR, podRN)
	case 1:
		return &list.Items[0], nil
	default:
		return nil, fmt.Errorf("multiple resources of type {%s} found for label selector {%s} in namespace {%s},"+
			" when only one was expected", podGR.String(), lSelector.String(), ns)
	}
}
