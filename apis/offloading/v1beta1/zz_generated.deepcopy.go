//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	corev1beta1 "github.com/liqotech/liqo/apis/core/v1beta1"
	"k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Affinity) DeepCopyInto(out *Affinity) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Affinity.
func (in *Affinity) DeepCopy() *Affinity {
	if in == nil {
		return nil
	}
	out := new(Affinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentTemplate) DeepCopyInto(out *DeploymentTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentTemplate.
func (in *DeploymentTemplate) DeepCopy() *DeploymentTemplate {
	if in == nil {
		return nil
	}
	out := new(DeploymentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSliceTemplate) DeepCopyInto(out *EndpointSliceTemplate) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]discoveryv1.Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]discoveryv1.EndpointPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSliceTemplate.
func (in *EndpointSliceTemplate) DeepCopy() *EndpointSliceTemplate {
	if in == nil {
		return nil
	}
	out := new(EndpointSliceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceMap) DeepCopyInto(out *NamespaceMap) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceMap.
func (in *NamespaceMap) DeepCopy() *NamespaceMap {
	if in == nil {
		return nil
	}
	out := new(NamespaceMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceMap) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceMapList) DeepCopyInto(out *NamespaceMapList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespaceMap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceMapList.
func (in *NamespaceMapList) DeepCopy() *NamespaceMapList {
	if in == nil {
		return nil
	}
	out := new(NamespaceMapList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceMapList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceMapSpec) DeepCopyInto(out *NamespaceMapSpec) {
	*out = *in
	if in.DesiredMapping != nil {
		in, out := &in.DesiredMapping, &out.DesiredMapping
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceMapSpec.
func (in *NamespaceMapSpec) DeepCopy() *NamespaceMapSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceMapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceMapStatus) DeepCopyInto(out *NamespaceMapStatus) {
	*out = *in
	if in.CurrentMapping != nil {
		in, out := &in.CurrentMapping, &out.CurrentMapping
		*out = make(map[string]RemoteNamespaceStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceMapStatus.
func (in *NamespaceMapStatus) DeepCopy() *NamespaceMapStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceMapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceOffloading) DeepCopyInto(out *NamespaceOffloading) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceOffloading.
func (in *NamespaceOffloading) DeepCopy() *NamespaceOffloading {
	if in == nil {
		return nil
	}
	out := new(NamespaceOffloading)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceOffloading) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceOffloadingList) DeepCopyInto(out *NamespaceOffloadingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespaceOffloading, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceOffloadingList.
func (in *NamespaceOffloadingList) DeepCopy() *NamespaceOffloadingList {
	if in == nil {
		return nil
	}
	out := new(NamespaceOffloadingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceOffloadingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceOffloadingSpec) DeepCopyInto(out *NamespaceOffloadingSpec) {
	*out = *in
	in.ClusterSelector.DeepCopyInto(&out.ClusterSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceOffloadingSpec.
func (in *NamespaceOffloadingSpec) DeepCopy() *NamespaceOffloadingSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceOffloadingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceOffloadingStatus) DeepCopyInto(out *NamespaceOffloadingStatus) {
	*out = *in
	if in.RemoteNamespacesConditions != nil {
		in, out := &in.RemoteNamespacesConditions, &out.RemoteNamespacesConditions
		*out = make(map[string]RemoteNamespaceConditions, len(*in))
		for key, val := range *in {
			var outVal []RemoteNamespaceCondition
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(RemoteNamespaceConditions, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceOffloadingStatus.
func (in *NamespaceOffloadingStatus) DeepCopy() *NamespaceOffloadingStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceOffloadingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OffloadingPatch) DeepCopyInto(out *OffloadingPatch) {
	*out = *in
	if in.AnnotationsNotReflected != nil {
		in, out := &in.AnnotationsNotReflected, &out.AnnotationsNotReflected
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LabelsNotReflected != nil {
		in, out := &in.LabelsNotReflected, &out.LabelsNotReflected
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OffloadingPatch.
func (in *OffloadingPatch) DeepCopy() *OffloadingPatch {
	if in == nil {
		return nil
	}
	out := new(OffloadingPatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Quota) DeepCopyInto(out *Quota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Quota.
func (in *Quota) DeepCopy() *Quota {
	if in == nil {
		return nil
	}
	out := new(Quota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Quota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaList) DeepCopyInto(out *QuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Quota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaList.
func (in *QuotaList) DeepCopy() *QuotaList {
	if in == nil {
		return nil
	}
	out := new(QuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaSpec) DeepCopyInto(out *QuotaSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Cordoned != nil {
		in, out := &in.Cordoned, &out.Cordoned
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaSpec.
func (in *QuotaSpec) DeepCopy() *QuotaSpec {
	if in == nil {
		return nil
	}
	out := new(QuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReflectorConfig) DeepCopyInto(out *ReflectorConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReflectorConfig.
func (in *ReflectorConfig) DeepCopy() *ReflectorConfig {
	if in == nil {
		return nil
	}
	out := new(ReflectorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteNamespaceCondition) DeepCopyInto(out *RemoteNamespaceCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteNamespaceCondition.
func (in *RemoteNamespaceCondition) DeepCopy() *RemoteNamespaceCondition {
	if in == nil {
		return nil
	}
	out := new(RemoteNamespaceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RemoteNamespaceConditions) DeepCopyInto(out *RemoteNamespaceConditions) {
	{
		in := &in
		*out = make(RemoteNamespaceConditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteNamespaceConditions.
func (in RemoteNamespaceConditions) DeepCopy() RemoteNamespaceConditions {
	if in == nil {
		return nil
	}
	out := new(RemoteNamespaceConditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteNamespaceStatus) DeepCopyInto(out *RemoteNamespaceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteNamespaceStatus.
func (in *RemoteNamespaceStatus) DeepCopy() *RemoteNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(RemoteNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShadowEndpointSlice) DeepCopyInto(out *ShadowEndpointSlice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShadowEndpointSlice.
func (in *ShadowEndpointSlice) DeepCopy() *ShadowEndpointSlice {
	if in == nil {
		return nil
	}
	out := new(ShadowEndpointSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShadowEndpointSlice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShadowEndpointSliceList) DeepCopyInto(out *ShadowEndpointSliceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ShadowEndpointSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShadowEndpointSliceList.
func (in *ShadowEndpointSliceList) DeepCopy() *ShadowEndpointSliceList {
	if in == nil {
		return nil
	}
	out := new(ShadowEndpointSliceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShadowEndpointSliceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShadowEndpointSliceSpec) DeepCopyInto(out *ShadowEndpointSliceSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShadowEndpointSliceSpec.
func (in *ShadowEndpointSliceSpec) DeepCopy() *ShadowEndpointSliceSpec {
	if in == nil {
		return nil
	}
	out := new(ShadowEndpointSliceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShadowPod) DeepCopyInto(out *ShadowPod) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShadowPod.
func (in *ShadowPod) DeepCopy() *ShadowPod {
	if in == nil {
		return nil
	}
	out := new(ShadowPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShadowPod) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShadowPodList) DeepCopyInto(out *ShadowPodList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ShadowPod, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShadowPodList.
func (in *ShadowPodList) DeepCopy() *ShadowPodList {
	if in == nil {
		return nil
	}
	out := new(ShadowPodList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShadowPodList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShadowPodSpec) DeepCopyInto(out *ShadowPodSpec) {
	*out = *in
	in.Pod.DeepCopyInto(&out.Pod)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShadowPodSpec.
func (in *ShadowPodSpec) DeepCopy() *ShadowPodSpec {
	if in == nil {
		return nil
	}
	out := new(ShadowPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShadowPodStatus) DeepCopyInto(out *ShadowPodStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShadowPodStatus.
func (in *ShadowPodStatus) DeepCopy() *ShadowPodStatus {
	if in == nil {
		return nil
	}
	out := new(ShadowPodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNode) DeepCopyInto(out *VirtualNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNode.
func (in *VirtualNode) DeepCopy() *VirtualNode {
	if in == nil {
		return nil
	}
	out := new(VirtualNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNodeCondition) DeepCopyInto(out *VirtualNodeCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNodeCondition.
func (in *VirtualNodeCondition) DeepCopy() *VirtualNodeCondition {
	if in == nil {
		return nil
	}
	out := new(VirtualNodeCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNodeList) DeepCopyInto(out *VirtualNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNodeList.
func (in *VirtualNodeList) DeepCopy() *VirtualNodeList {
	if in == nil {
		return nil
	}
	out := new(VirtualNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNodeSpec) DeepCopyInto(out *VirtualNodeSpec) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(DeploymentTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.OffloadingPatch != nil {
		in, out := &in.OffloadingPatch, &out.OffloadingPatch
		*out = new(OffloadingPatch)
		(*in).DeepCopyInto(*out)
	}
	if in.CreateNode != nil {
		in, out := &in.CreateNode, &out.CreateNode
		*out = new(bool)
		**out = **in
	}
	if in.DisableNetworkCheck != nil {
		in, out := &in.DisableNetworkCheck, &out.DisableNetworkCheck
		*out = new(bool)
		**out = **in
	}
	if in.KubeconfigSecretRef != nil {
		in, out := &in.KubeconfigSecretRef, &out.KubeconfigSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]v1.ContainerImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ResourceQuota.DeepCopyInto(&out.ResourceQuota)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageClasses != nil {
		in, out := &in.StorageClasses, &out.StorageClasses
		*out = make([]corev1beta1.StorageType, len(*in))
		copy(*out, *in)
	}
	if in.IngressClasses != nil {
		in, out := &in.IngressClasses, &out.IngressClasses
		*out = make([]corev1beta1.IngressType, len(*in))
		copy(*out, *in)
	}
	if in.LoadBalancerClasses != nil {
		in, out := &in.LoadBalancerClasses, &out.LoadBalancerClasses
		*out = make([]corev1beta1.LoadBalancerType, len(*in))
		copy(*out, *in)
	}
	if in.VkOptionsTemplateRef != nil {
		in, out := &in.VkOptionsTemplateRef, &out.VkOptionsTemplateRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNodeSpec.
func (in *VirtualNodeSpec) DeepCopy() *VirtualNodeSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNodeStatus) DeepCopyInto(out *VirtualNodeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]VirtualNodeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNodeStatus.
func (in *VirtualNodeStatus) DeepCopy() *VirtualNodeStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VkOptionsTemplate) DeepCopyInto(out *VkOptionsTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VkOptionsTemplate.
func (in *VkOptionsTemplate) DeepCopy() *VkOptionsTemplate {
	if in == nil {
		return nil
	}
	out := new(VkOptionsTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VkOptionsTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VkOptionsTemplateList) DeepCopyInto(out *VkOptionsTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VkOptionsTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VkOptionsTemplateList.
func (in *VkOptionsTemplateList) DeepCopy() *VkOptionsTemplateList {
	if in == nil {
		return nil
	}
	out := new(VkOptionsTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VkOptionsTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VkOptionsTemplateSpec) DeepCopyInto(out *VkOptionsTemplateSpec) {
	*out = *in
	if in.LabelsNotReflected != nil {
		in, out := &in.LabelsNotReflected, &out.LabelsNotReflected
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AnnotationsNotReflected != nil {
		in, out := &in.AnnotationsNotReflected, &out.AnnotationsNotReflected
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReflectorsConfig != nil {
		in, out := &in.ReflectorsConfig, &out.ReflectorsConfig
		*out = make(map[string]ReflectorConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraAnnotations != nil {
		in, out := &in.ExtraAnnotations, &out.ExtraAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExtraLabels != nil {
		in, out := &in.ExtraLabels, &out.ExtraLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeExtraAnnotations != nil {
		in, out := &in.NodeExtraAnnotations, &out.NodeExtraAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeExtraLabels != nil {
		in, out := &in.NodeExtraLabels, &out.NodeExtraLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VkOptionsTemplateSpec.
func (in *VkOptionsTemplateSpec) DeepCopy() *VkOptionsTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(VkOptionsTemplateSpec)
	in.DeepCopyInto(out)
	return out
}
