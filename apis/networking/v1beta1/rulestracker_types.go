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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RulesTrackerStatus defines the observed state of RulesTracker.
type RulesTrackerStatus struct {
	// RouteConfigurations tracks the status of RouteConfiguration resources.
	// The key is the name of the RouteConfiguration resource.
	// The value is an array of standard Kubernetes metav1.Condition objects.
	// +optional
	RouteConfigurations map[string][]metav1.Condition `json:"routeConfigurations,omitempty"`

	// FirewallConfigurations tracks the status of FirewallConfiguration resources.
	// The key is the name of the FirewallConfiguration resource.
	// The value is an array of standard Kubernetes metav1.Condition objects.
	// +optional
	FirewallConfigurations map[string][]metav1.Condition `json:"firewallConfigurations,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=liqo,shortName=rt;rtrack
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// RulesTracker tracks the status of RouteConfiguration and FirewallConfiguration resources.
type RulesTracker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status RulesTrackerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RulesTrackerList contains a list of RulesTracker.
type RulesTrackerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RulesTracker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RulesTracker{}, &RulesTrackerList{})
}
