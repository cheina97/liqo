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
	"fmt"

	"k8s.io/apimachinery/pkg/labels"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/liqotech/liqo/pkg/liqoctl/common"
	"github.com/liqotech/liqo/pkg/utils"
	"github.com/liqotech/liqo/pkg/utils/getters"
)

// LocalInfoChecker implements the Check interface.
// holds the Localinformation about the cluster.
type LocalInfoChecker struct {
	client            client.Client
	namespace         string
	rootLocalInfoNode InfoNode
	collectionErrors  []collectionError
}

// newPodChecker return a new pod checker.
func newLocalInfoChecker(namespace string, cl client.Client) *LocalInfoChecker {
	return &LocalInfoChecker{
		client:            cl,
		namespace:         namespace,
		rootLocalInfoNode: newRootInfoNode("Local Cluster Informations"),
	}
}

// Collect implements the collect method of the Checker interface.
// it collects the infos of the local cluster.
func (lic *LocalInfoChecker) Collect(ctx context.Context) error {
	clusterIdentity, err := utils.GetClusterIdentityWithControllerClient(ctx, lic.client, lic.namespace)
	if err != nil {
		lic.addCollectionError("ClusterIdentity", "", err)
	}
	clusterIdentityNode := lic.rootLocalInfoNode.addSectionToNode("Cluster Identity", "")
	clusterIdentityNode.addDataToNode("Cluster ID", clusterIdentity.ClusterID)
	clusterIdentityNode.addDataToNode("Cluster Name", clusterIdentity.ClusterName)

	clusterLabelsNode := lic.rootLocalInfoNode.addSectionToNode("Cluster Labels", "")
	hc, err := common.NewLiqoHelmClient()
	if err != nil {
		lic.addCollectionError("ClusterIdentity", "Creating new LiqoHelmClient", err)
	}
	clusterLabels, err := hc.GetClusterLabels()
	if err != nil {
		lic.addCollectionError("Cluster Labels", "Getting cluster labels", err)
	}
	for k, v := range clusterLabels {
		clusterLabelsNode.addDataToNode(k, v)
	}

	ipamNode := lic.rootLocalInfoNode.addSectionToNode("IPAM", "")
	ipamStorage, err := getters.GetIPAMStorageByLabel(ctx, lic.client, "default", labels.NewSelector())
	if err != nil {
		lic.addCollectionError("IPAM", "", err)
	}
	ipamNode.addDataToNode("Pod CIDR", ipamStorage.Spec.PodCIDR)
	ipamNode.addDataToNode("External CIDR", ipamStorage.Spec.ExternalCIDR)
	ipamNode.addDataToNode("Service CIDR", ipamStorage.Spec.ServiceCIDR)
	if len(ipamStorage.Spec.ReservedSubnets) != 0 {
		ipamNode.addDataListToNode("Reserved Subnets", ipamStorage.Spec.ReservedSubnets)
	}

	return nil
}

// Format implements the format method of the Checker interface.
// it outputs the infos about the local cluster in a string ready to be
// printed out.
func (lic *LocalInfoChecker) Format() (string, error) {
	w, buf := newTabWriter("")

	fmt.Fprintf(w, "%s", deepPrintInfo(&lic.rootLocalInfoNode))

	for _, err := range lic.collectionErrors {
		fmt.Fprintf(w, "%s\t%s\t%s%s%s\n", err.appType, err.appName, red, err.err, reset)
	}

	if err := w.Flush(); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// HasSucceeded return true if no errors have been kept.
func (lic *LocalInfoChecker) HasSucceeded() bool {
	if len(lic.collectionErrors) == 0 {
		return true
	}
	return false
}

// addCollectionError adds a collection error. A collection error is an error that happens while
// collecting the status of a Liqo component.
func (lic *LocalInfoChecker) addCollectionError(localInfoType, localInfoName string, err error) {
	lic.collectionErrors = append(lic.collectionErrors, newCollectionError(localInfoType, localInfoName, err))
}
