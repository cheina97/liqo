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

package autocompletion

import (
	"context"

	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/liqotech/liqo/apis/discovery/v1alpha1"
	"github.com/liqotech/liqo/pkg/liqoctl/common"
	"github.com/liqotech/liqo/pkg/utils/getters"
)

// GetClusters returns the list of foreign clusters that are peered with the local cluster.
func GetClusters(ctx context.Context) (*v1alpha1.ForeignClusterList, error) {
	restConfig, err := common.GetLiqoctlRestConf()
	if err != nil {
		return nil, err
	}

	cl, err := client.New(restConfig, client.Options{})
	if err != nil {
		return nil, err
	}

	foreignClusters, err := getters.GetForeignClustersByLabel(ctx, cl, labels.NewSelector())
	if err != nil {
		return nil, err
	}
	return foreignClusters, nil
}

// GetClusterNames returns the list of foreign cluster names that start with the given string.
func GetClusterNames(ctx context.Context) ([]string, error) {
	var clustersName []string
	foreignClusters, err := GetClusters(ctx)
	if err != nil {
		return nil, err
	}
	for i := range foreignClusters.Items {
		clustersName = append(clustersName, foreignClusters.Items[i].Spec.ClusterIdentity.ClusterName)
	}
	return clustersName, nil
}

// GetClusterIDs returns the list of foreignClusters cluster IDs that are peered with the local cluster.
func GetClusterIDs(ctx context.Context) ([]string, error) {
	var clustersID []string
	foreignClusters, err := GetClusters(ctx)
	if err != nil {
		return nil, err
	}
	for i := range foreignClusters.Items {
		clustersID = append(clustersID, foreignClusters.Items[i].Spec.ClusterIdentity.ClusterID)
	}
	return clustersID, nil
}
