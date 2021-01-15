// +build !ignore_autogenerated

// Copyright 2021 Vectorized, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import runtime "k8s.io/apimachinery/pkg/runtime"

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedpandaCluster) DeepCopyInto(out *RedpandaCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedpandaCluster.
func (in *RedpandaCluster) DeepCopy() *RedpandaCluster {
	if in == nil {
		return nil
	}
	out := new(RedpandaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedpandaCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedpandaClusterList) DeepCopyInto(out *RedpandaClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedpandaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedpandaClusterList.
func (in *RedpandaClusterList) DeepCopy() *RedpandaClusterList {
	if in == nil {
		return nil
	}
	out := new(RedpandaClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedpandaClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedpandaClusterSpec) DeepCopyInto(out *RedpandaClusterSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedpandaClusterSpec.
func (in *RedpandaClusterSpec) DeepCopy() *RedpandaClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RedpandaClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedpandaClusterStatus) DeepCopyInto(out *RedpandaClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedpandaClusterStatus.
func (in *RedpandaClusterStatus) DeepCopy() *RedpandaClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RedpandaClusterStatus)
	in.DeepCopyInto(out)
	return out
}
