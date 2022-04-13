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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/liqotech/liqo/pkg/consts"
	"github.com/liqotech/liqo/pkg/utils"
	"github.com/liqotech/liqo/pkg/utils/getters"
	"github.com/liqotech/liqo/pkg/utils/slice"
)

// RemoteInfoChecker implements the Check interface.
// holds informations about remote clusters.
type RemoteInfoChecker struct {
	client             client.Client
	namespace          string
	clusterNameFilter  []string
	clusterIDFilter    []string
	errors             bool
	rootRemoteInfoNode InfoNode
	collectionErrors   []collectionError
}

// newRemoteInfoChecker return a new remote info checker.
func newRemoteInfoChecker(namespace string, clusterNameFilter, clusterIDFilter []string, cl client.Client) *RemoteInfoChecker {
	return &RemoteInfoChecker{
		client:             cl,
		namespace:          namespace,
		clusterNameFilter:  clusterNameFilter,
		clusterIDFilter:    clusterIDFilter,
		errors:             false,
		rootRemoteInfoNode: newRootInfoNode("Remote Clusters Information"),
	}
}

// argsFilterCheck returns true if the given foreignClusterID or foreignClusterName are contained
// in the respective filter list or if filter lists are void.
func (ric *RemoteInfoChecker) argsFilterCheck(foreignClusterID, foreignClusterName string) bool {
	return slice.ContainsString(ric.clusterIDFilter, foreignClusterID) ||
		slice.ContainsString(ric.clusterNameFilter, foreignClusterName) ||
		(len(ric.clusterIDFilter) == 0 && len(ric.clusterNameFilter) == 0)
}

// Collect implements the collect method of the Checker interface.
// it collects the infos of the Remote cluster.
func (ric *RemoteInfoChecker) Collect(ctx context.Context) error {
	var localNetworkConfigNode, remoteNetworkConfigNode, selectedNode *InfoNode
	var remoteNodeMsg, remotePodCIDRMsg, remoteExternalCIDRMsg string
	localClusterIdentity, err := utils.GetClusterIdentityWithControllerClient(ctx, ric.client, ric.namespace)
	localClusterName := localClusterIdentity.ClusterName
	if err != nil {
		ric.addCollectionError("LocalClusterName", "", err)
		ric.errors = true
	}

	foreignClusters, err := getters.GetForeignClustersByLabel(ctx, ric.client, labels.NewSelector())
	if err == nil {
		for i := range foreignClusters.Items {
			foreignClusterName := foreignClusters.Items[i].Spec.ClusterIdentity.ClusterName
			foreignClusterID := foreignClusters.Items[i].Spec.ClusterIdentity.ClusterID
			foreignClusterTenantNamespace := foreignClusters.Items[i].Status.TenantNamespace

			if ric.argsFilterCheck(foreignClusterID, foreignClusterName) {
				networkConfigs, err := getters.GetNetworkConfigsByLabel(ctx, ric.client, foreignClusterTenantNamespace.Local, labels.NewSelector())
				if err != nil {
					ric.addCollectionError("NetworkConfigs", "unable to collect NetworkConfigs", err)
					ric.errors = true
				}

				clusterNode := ric.rootRemoteInfoNode.addSectionToNode(foreignClusterName, "")
				localNetworkConfigNode = clusterNode.addSectionToNode("Local Network Configuration", "")
				remoteNetworkConfigNode = clusterNode.addSectionToNode("Remote Network Configuration", "")
				for i := range networkConfigs.Items {
					if networkConfigs.Items[i].Labels[consts.ReplicationOriginLabel] == "true" {
						selectedNode = localNetworkConfigNode
						remoteNodeMsg = fmt.Sprintf("Status: how %s's CIDRs has been remapped by %s", localClusterName, foreignClusterName)
					} else {
						selectedNode = remoteNetworkConfigNode
						remoteNodeMsg = fmt.Sprintf("Status: how %s remapped %s's CIDRs", localClusterName, foreignClusterName)
					}

					if networkConfigs.Items[i].Status.PodCIDRNAT == consts.DefaultCIDRValue {
						remotePodCIDRMsg = NotRemappedMsg
					} else {
						remotePodCIDRMsg = networkConfigs.Items[i].Status.PodCIDRNAT
					}
					if networkConfigs.Items[i].Status.ExternalCIDRNAT == consts.DefaultCIDRValue {
						remoteExternalCIDRMsg = NotRemappedMsg
					} else {
						remoteExternalCIDRMsg = networkConfigs.Items[i].Status.ExternalCIDRNAT
					}

					originalNode := selectedNode.addSectionToNode("Original Network Configuration", "Spec")
					originalNode.addDataToNode("Pod CIDR", networkConfigs.Items[i].Spec.PodCIDR)
					originalNode.addDataToNode("External CIDR", networkConfigs.Items[i].Spec.ExternalCIDR)
					remoteNode := selectedNode.addSectionToNode("Remapped Network Configuration", remoteNodeMsg)
					remoteNode.addDataToNode("Pod CIDR", remotePodCIDRMsg)
					remoteNode.addDataToNode("External CIDR", remoteExternalCIDRMsg)
				}
				tunnelEnpointSelector, err := metav1.LabelSelectorAsSelector(&metav1.LabelSelector{
					MatchExpressions: []metav1.LabelSelectorRequirement{
						{
							Key:      "clusterID",
							Operator: metav1.LabelSelectorOpIn,
							Values:   []string{foreignClusterID},
						},
					},
				})
				if err != nil {
					ric.addCollectionError("TunnelEndpoint", fmt.Sprintf("unable to create TunnelEnpoint Selector for %s cluster", foreignClusterID), err)
					ric.errors = true
				}

				tunnelEndpoint, err := getters.GetTunnelEndpointByLabel(ctx, ric.client, "", tunnelEnpointSelector)
				if err != nil {
					ric.addCollectionError("TunnelEndpoint", fmt.Sprintf("unable to get TunnelEndpoint for %s cluster", foreignClusterID), err)
					ric.errors = true
				}
				tunnelEndpointNode := clusterNode.addSectionToNode("Tunnel Endpoint", "")
				tunnelEndpointNode.addDataToNode("Gateway IP", tunnelEndpoint.Status.GatewayIP)
				tunnelEndpointNode.addDataToNode("Endpoint IP", tunnelEndpoint.Status.Connection.PeerConfiguration["endpointIP"])
			}
		}
	}
	if len(ric.rootRemoteInfoNode.nextNodes) == 0 {
		ric.rootRemoteInfoNode.addSectionToNode("Remote clusters not found ...", "")
	}

	return nil
}

// Format implements the format method of the Checker interface.
// it outputs the infos about the Remote cluster in a string ready to be
// printed out.
func (ric *RemoteInfoChecker) Format() (string, error) {
	w, buf := newTabWriter("")

	fmt.Fprintf(w, "%s", deepPrintInfo(&ric.rootRemoteInfoNode))

	for _, err := range ric.collectionErrors {
		fmt.Fprintf(w, "%s\t%s\t%s%s%s\n", err.appType, err.appName, red, err.err, reset)
	}

	if err := w.Flush(); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// HasSucceeded return true if no errors have been kept.
func (ric *RemoteInfoChecker) HasSucceeded() bool {
	return !ric.errors
}

// addCollectionError adds a collection error. A collection error is an error that happens while
// collecting the status of a Liqo component.
func (ric *RemoteInfoChecker) addCollectionError(remoteInfoType, remoteInfoName string, err error) {
	ric.collectionErrors = append(ric.collectionErrors, newCollectionError(remoteInfoType, remoteInfoName, err))
}
