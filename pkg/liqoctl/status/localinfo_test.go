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

package status

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/liqotech/liqo/apis/discovery/v1alpha1"
	netv1alpha1 "github.com/liqotech/liqo/apis/net/v1alpha1"
	"github.com/liqotech/liqo/pkg/utils"
)

var _ = Describe("LocalInfo", func() {
	const (
		clusterID   = "fake"
		clusterName = "fake"
		namespace   = "liqo"
		rootTitle   = "Local Cluster Informations"
	)

	var (
		clientBuilder   fake.ClientBuilder
		rootNode        = newRootInfoNode(rootTitle)
		lic             *LocalInfoChecker
		ctx             = context.Background()
		clusterIdentity v1alpha1.ClusterIdentity
	)

	BeforeEach(func() {
		ctx = context.Background()
		_ = netv1alpha1.AddToScheme(scheme.Scheme)
		clientBuilder = *fake.NewClientBuilder().WithScheme(scheme.Scheme)
	})

	Context("Creating a new LocalInfoChecker", func() {
		JustBeforeEach(func() {
			lic = newLocalInfoChecker(namespace, clientBuilder.Build())
		})
		It("should return a valid LocalInfoChecker", func() {
			licTest := &LocalInfoChecker{
				client:            clientBuilder.Build(),
				namespace:         namespace,
				collectionErrors:  nil,
				rootLocalInfoNode: rootNode,
			}
			Expect(*lic).To(Equal(*licTest))
		})
	})
	Context("Getting a local cluster identity", func() {
		BeforeEach(func() {
			clientBuilder.WithObjects(&corev1.ConfigMap{
				TypeMeta: metav1.TypeMeta{
					Kind: "ConfigMap",
				},
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app.kubernetes.io/name": "clusterid-configmap",
					},
					Namespace: namespace,
				},
				Immutable: new(bool),
				Data: map[string]string{
					"CLUSTER_ID":   clusterID,
					"CLUSTER_NAME": clusterName,
				},
				BinaryData: map[string][]byte{},
			})
		})
		JustBeforeEach(func() {
			clientBuilder.Build()
			clusterIdentity, _ = utils.GetClusterIdentityWithControllerClient(ctx, clientBuilder.Build(), namespace)
		})
		It("should return a valid cluster identity", func() {
			Expect(clusterIdentity.ClusterID).To(Equal(clusterID))
			Expect(clusterIdentity.ClusterName).To(Equal(clusterName))
		})
	})
})
