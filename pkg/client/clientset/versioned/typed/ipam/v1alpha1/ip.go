// Copyright 2019-2025 The Liqo Authors
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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"

	v1alpha1 "github.com/liqotech/liqo/apis/ipam/v1alpha1"
	scheme "github.com/liqotech/liqo/pkg/client/clientset/versioned/scheme"
)

// IPsGetter has a method to return a IPInterface.
// A group's client should implement this interface.
type IPsGetter interface {
	IPs(namespace string) IPInterface
}

// IPInterface has methods to work with IP resources.
type IPInterface interface {
	Create(ctx context.Context, iP *v1alpha1.IP, opts v1.CreateOptions) (*v1alpha1.IP, error)
	Update(ctx context.Context, iP *v1alpha1.IP, opts v1.UpdateOptions) (*v1alpha1.IP, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, iP *v1alpha1.IP, opts v1.UpdateOptions) (*v1alpha1.IP, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.IP, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.IPList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IP, err error)
	IPExpansion
}

// iPs implements IPInterface
type iPs struct {
	*gentype.ClientWithList[*v1alpha1.IP, *v1alpha1.IPList]
}

// newIPs returns a IPs
func newIPs(c *IpamV1alpha1Client, namespace string) *iPs {
	return &iPs{
		gentype.NewClientWithList[*v1alpha1.IP, *v1alpha1.IPList](
			"ips",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.IP { return &v1alpha1.IP{} },
			func() *v1alpha1.IPList { return &v1alpha1.IPList{} }),
	}
}
