// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

package testing

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExternalNamespacedType).DeepCopyInto(out.(*ExternalNamespacedType))
			return nil
		}, InType: reflect.TypeOf(&ExternalNamespacedType{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExternalNamespacedType2).DeepCopyInto(out.(*ExternalNamespacedType2))
			return nil
		}, InType: reflect.TypeOf(&ExternalNamespacedType2{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExternalType).DeepCopyInto(out.(*ExternalType))
			return nil
		}, InType: reflect.TypeOf(&ExternalType{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExternalType2).DeepCopyInto(out.(*ExternalType2))
			return nil
		}, InType: reflect.TypeOf(&ExternalType2{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InternalNamespacedType).DeepCopyInto(out.(*InternalNamespacedType))
			return nil
		}, InType: reflect.TypeOf(&InternalNamespacedType{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InternalType).DeepCopyInto(out.(*InternalType))
			return nil
		}, InType: reflect.TypeOf(&InternalType{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNamespacedType) DeepCopyInto(out *ExternalNamespacedType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNamespacedType.
func (in *ExternalNamespacedType) DeepCopy() *ExternalNamespacedType {
	if in == nil {
		return nil
	}
	out := new(ExternalNamespacedType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalNamespacedType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNamespacedType2) DeepCopyInto(out *ExternalNamespacedType2) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNamespacedType2.
func (in *ExternalNamespacedType2) DeepCopy() *ExternalNamespacedType2 {
	if in == nil {
		return nil
	}
	out := new(ExternalNamespacedType2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalNamespacedType2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalType) DeepCopyInto(out *ExternalType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalType.
func (in *ExternalType) DeepCopy() *ExternalType {
	if in == nil {
		return nil
	}
	out := new(ExternalType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalType2) DeepCopyInto(out *ExternalType2) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalType2.
func (in *ExternalType2) DeepCopy() *ExternalType2 {
	if in == nil {
		return nil
	}
	out := new(ExternalType2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalType2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalNamespacedType) DeepCopyInto(out *InternalNamespacedType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalNamespacedType.
func (in *InternalNamespacedType) DeepCopy() *InternalNamespacedType {
	if in == nil {
		return nil
	}
	out := new(InternalNamespacedType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InternalNamespacedType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalType) DeepCopyInto(out *InternalType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalType.
func (in *InternalType) DeepCopy() *InternalType {
	if in == nil {
		return nil
	}
	out := new(InternalType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InternalType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}