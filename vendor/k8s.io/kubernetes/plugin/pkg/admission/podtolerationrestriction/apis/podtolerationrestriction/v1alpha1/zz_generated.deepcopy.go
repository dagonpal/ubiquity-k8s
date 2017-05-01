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

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Configuration, InType: reflect.TypeOf(&Configuration{})},
	)
}

func DeepCopy_v1alpha1_Configuration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Configuration)
		out := out.(*Configuration)
		*out = *in
		if in.Default != nil {
			in, out := &in.Default, &out.Default
			*out = make([]v1.Toleration, len(*in))
			for i := range *in {
				if err := v1.DeepCopy_v1_Toleration(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Whitelist != nil {
			in, out := &in.Whitelist, &out.Whitelist
			*out = make([]v1.Toleration, len(*in))
			for i := range *in {
				if err := v1.DeepCopy_v1_Toleration(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}
