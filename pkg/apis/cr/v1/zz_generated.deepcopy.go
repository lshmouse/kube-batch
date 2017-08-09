// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	reflect "reflect"
)

// Deprecated: GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ResourceList).DeepCopyInto(out.(*ResourceList))
			return nil
		}, InType: reflect.TypeOf(&ResourceList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ResourceQuotaAllocator).DeepCopyInto(out.(*ResourceQuotaAllocator))
			return nil
		}, InType: reflect.TypeOf(&ResourceQuotaAllocator{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ResourceQuotaAllocatorList).DeepCopyInto(out.(*ResourceQuotaAllocatorList))
			return nil
		}, InType: reflect.TypeOf(&ResourceQuotaAllocatorList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ResourceQuotaAllocatorSpec).DeepCopyInto(out.(*ResourceQuotaAllocatorSpec))
			return nil
		}, InType: reflect.TypeOf(&ResourceQuotaAllocatorSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ResourceQuotaAllocatorStatus).DeepCopyInto(out.(*ResourceQuotaAllocatorStatus))
			return nil
		}, InType: reflect.TypeOf(&ResourceQuotaAllocatorStatus{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceList) DeepCopyInto(out *ResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[ResourceName]resource.Quantity, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ResourceList.
func (x *ResourceList) DeepCopy() *ResourceList {
	if x == nil {
		return nil
	}
	out := new(ResourceList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ResourceList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaAllocator) DeepCopyInto(out *ResourceQuotaAllocator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaAllocator.
func (x *ResourceQuotaAllocator) DeepCopy() *ResourceQuotaAllocator {
	if x == nil {
		return nil
	}
	out := new(ResourceQuotaAllocator)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ResourceQuotaAllocator) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaAllocatorList) DeepCopyInto(out *ResourceQuotaAllocatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceQuotaAllocator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaAllocatorList.
func (x *ResourceQuotaAllocatorList) DeepCopy() *ResourceQuotaAllocatorList {
	if x == nil {
		return nil
	}
	out := new(ResourceQuotaAllocatorList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ResourceQuotaAllocatorList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaAllocatorSpec) DeepCopyInto(out *ResourceQuotaAllocatorSpec) {
	*out = *in
	if in.Share != nil {
		in, out := &in.Share, &out.Share
		*out = make(map[string]intstr.IntOrString, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaAllocatorSpec.
func (x *ResourceQuotaAllocatorSpec) DeepCopy() *ResourceQuotaAllocatorSpec {
	if x == nil {
		return nil
	}
	out := new(ResourceQuotaAllocatorSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaAllocatorStatus) DeepCopyInto(out *ResourceQuotaAllocatorStatus) {
	*out = *in
	if in.Share != nil {
		in, out := &in.Share, &out.Share
		*out = make(map[string]intstr.IntOrString, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Usage != nil {
		in, out := &in.Usage, &out.Usage
		*out = make(map[string]ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaAllocatorStatus.
func (x *ResourceQuotaAllocatorStatus) DeepCopy() *ResourceQuotaAllocatorStatus {
	if x == nil {
		return nil
	}
	out := new(ResourceQuotaAllocatorStatus)
	x.DeepCopyInto(out)
	return out
}
