//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"

	"go.goms.io/fleet/apis/placement/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceOverride) DeepCopyInto(out *ClusterResourceOverride) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceOverride.
func (in *ClusterResourceOverride) DeepCopy() *ClusterResourceOverride {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceOverride) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceOverrideList) DeepCopyInto(out *ClusterResourceOverrideList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterResourceOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceOverrideList.
func (in *ClusterResourceOverrideList) DeepCopy() *ClusterResourceOverrideList {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceOverrideList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceOverrideList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceOverrideSnapshot) DeepCopyInto(out *ClusterResourceOverrideSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceOverrideSnapshot.
func (in *ClusterResourceOverrideSnapshot) DeepCopy() *ClusterResourceOverrideSnapshot {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceOverrideSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceOverrideSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceOverrideSnapshotList) DeepCopyInto(out *ClusterResourceOverrideSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterResourceOverrideSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceOverrideSnapshotList.
func (in *ClusterResourceOverrideSnapshotList) DeepCopy() *ClusterResourceOverrideSnapshotList {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceOverrideSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceOverrideSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceOverrideSnapshotSpec) DeepCopyInto(out *ClusterResourceOverrideSnapshotSpec) {
	*out = *in
	in.OverrideSpec.DeepCopyInto(&out.OverrideSpec)
	if in.OverrideHash != nil {
		in, out := &in.OverrideHash, &out.OverrideHash
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceOverrideSnapshotSpec.
func (in *ClusterResourceOverrideSnapshotSpec) DeepCopy() *ClusterResourceOverrideSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceOverrideSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceOverrideSpec) DeepCopyInto(out *ClusterResourceOverrideSpec) {
	*out = *in
	if in.ClusterResourceSelectors != nil {
		in, out := &in.ClusterResourceSelectors, &out.ClusterResourceSelectors
		*out = make([]v1beta1.ClusterResourceSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(OverridePolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceOverrideSpec.
func (in *ClusterResourceOverrideSpec) DeepCopy() *ClusterResourceOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONPatchOverride) DeepCopyInto(out *JSONPatchOverride) {
	*out = *in
	in.Value.DeepCopyInto(&out.Value)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONPatchOverride.
func (in *JSONPatchOverride) DeepCopy() *JSONPatchOverride {
	if in == nil {
		return nil
	}
	out := new(JSONPatchOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OverridePolicy) DeepCopyInto(out *OverridePolicy) {
	*out = *in
	if in.OverrideRules != nil {
		in, out := &in.OverrideRules, &out.OverrideRules
		*out = make([]OverrideRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverridePolicy.
func (in *OverridePolicy) DeepCopy() *OverridePolicy {
	if in == nil {
		return nil
	}
	out := new(OverridePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OverrideRule) DeepCopyInto(out *OverrideRule) {
	*out = *in
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = new(v1beta1.ClusterSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.JSONPatchOverrides != nil {
		in, out := &in.JSONPatchOverrides, &out.JSONPatchOverrides
		*out = make([]JSONPatchOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverrideRule.
func (in *OverrideRule) DeepCopy() *OverrideRule {
	if in == nil {
		return nil
	}
	out := new(OverrideRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceOverride) DeepCopyInto(out *ResourceOverride) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceOverride.
func (in *ResourceOverride) DeepCopy() *ResourceOverride {
	if in == nil {
		return nil
	}
	out := new(ResourceOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceOverride) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceOverrideList) DeepCopyInto(out *ResourceOverrideList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceOverrideList.
func (in *ResourceOverrideList) DeepCopy() *ResourceOverrideList {
	if in == nil {
		return nil
	}
	out := new(ResourceOverrideList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceOverrideList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceOverrideSnapshot) DeepCopyInto(out *ResourceOverrideSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceOverrideSnapshot.
func (in *ResourceOverrideSnapshot) DeepCopy() *ResourceOverrideSnapshot {
	if in == nil {
		return nil
	}
	out := new(ResourceOverrideSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceOverrideSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceOverrideSnapshotList) DeepCopyInto(out *ResourceOverrideSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceOverrideSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceOverrideSnapshotList.
func (in *ResourceOverrideSnapshotList) DeepCopy() *ResourceOverrideSnapshotList {
	if in == nil {
		return nil
	}
	out := new(ResourceOverrideSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceOverrideSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceOverrideSnapshotSpec) DeepCopyInto(out *ResourceOverrideSnapshotSpec) {
	*out = *in
	in.OverrideSpec.DeepCopyInto(&out.OverrideSpec)
	if in.OverrideHash != nil {
		in, out := &in.OverrideHash, &out.OverrideHash
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceOverrideSnapshotSpec.
func (in *ResourceOverrideSnapshotSpec) DeepCopy() *ResourceOverrideSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceOverrideSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceOverrideSpec) DeepCopyInto(out *ResourceOverrideSpec) {
	*out = *in
	if in.ResourceSelectors != nil {
		in, out := &in.ResourceSelectors, &out.ResourceSelectors
		*out = make([]ResourceSelector, len(*in))
		copy(*out, *in)
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(OverridePolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceOverrideSpec.
func (in *ResourceOverrideSpec) DeepCopy() *ResourceOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSelector) DeepCopyInto(out *ResourceSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSelector.
func (in *ResourceSelector) DeepCopy() *ResourceSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceSelector)
	in.DeepCopyInto(out)
	return out
}