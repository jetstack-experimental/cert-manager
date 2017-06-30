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

package certmanager

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_certmanager_ACMECertData, InType: reflect.TypeOf(&ACMECertData{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_certmanager_ACMECertDetails, InType: reflect.TypeOf(&ACMECertDetails{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_certmanager_ACMEUserData, InType: reflect.TypeOf(&ACMEUserData{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_certmanager_Certificate, InType: reflect.TypeOf(&Certificate{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_certmanager_CertificateList, InType: reflect.TypeOf(&CertificateList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_certmanager_CertificateSpec, InType: reflect.TypeOf(&CertificateSpec{})},
	)
}

func DeepCopy_certmanager_ACMECertData(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ACMECertData)
		out := out.(*ACMECertData)
		*out = *in
		if in.Cert != nil {
			in, out := &in.Cert, &out.Cert
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		if in.PrivateKey != nil {
			in, out := &in.PrivateKey, &out.PrivateKey
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_certmanager_ACMECertDetails(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ACMECertDetails)
		out := out.(*ACMECertDetails)
		*out = *in
		return nil
	}
}

func DeepCopy_certmanager_ACMEUserData(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ACMEUserData)
		out := out.(*ACMEUserData)
		*out = *in
		if in.Key != nil {
			in, out := &in.Key, &out.Key
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_certmanager_Certificate(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Certificate)
		out := out.(*Certificate)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*CertificateSpec)
		}
		return nil
	}
}

func DeepCopy_certmanager_CertificateList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateList)
		out := out.(*CertificateList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Certificate, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Certificate)
				}
			}
		}
		return nil
	}
}

func DeepCopy_certmanager_CertificateSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSpec)
		out := out.(*CertificateSpec)
		*out = *in
		if in.AltNames != nil {
			in, out := &in.AltNames, &out.AltNames
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}
