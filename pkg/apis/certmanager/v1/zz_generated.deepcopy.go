// +build !ignore_autogenerated

/*
Copyright The cert-manager Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	acmev1 "github.com/cert-manager/cert-manager/pkg/apis/acme/v1"
	apismetav1 "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAIssuer) DeepCopyInto(out *CAIssuer) {
	*out = *in
	if in.CRLDistributionPoints != nil {
		in, out := &in.CRLDistributionPoints, &out.CRLDistributionPoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OCSPServers != nil {
		in, out := &in.OCSPServers, &out.OCSPServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAIssuer.
func (in *CAIssuer) DeepCopy() *CAIssuer {
	if in == nil {
		return nil
	}
	out := new(CAIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Certificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateCondition) DeepCopyInto(out *CertificateCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateCondition.
func (in *CertificateCondition) DeepCopy() *CertificateCondition {
	if in == nil {
		return nil
	}
	out := new(CertificateCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateKeystores) DeepCopyInto(out *CertificateKeystores) {
	*out = *in
	if in.JKS != nil {
		in, out := &in.JKS, &out.JKS
		*out = new(JKSKeystore)
		**out = **in
	}
	if in.PKCS12 != nil {
		in, out := &in.PKCS12, &out.PKCS12
		*out = new(PKCS12Keystore)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateKeystores.
func (in *CertificateKeystores) DeepCopy() *CertificateKeystores {
	if in == nil {
		return nil
	}
	out := new(CertificateKeystores)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateList) DeepCopyInto(out *CertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Certificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateList.
func (in *CertificateList) DeepCopy() *CertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatePrivateKey) DeepCopyInto(out *CertificatePrivateKey) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatePrivateKey.
func (in *CertificatePrivateKey) DeepCopy() *CertificatePrivateKey {
	if in == nil {
		return nil
	}
	out := new(CertificatePrivateKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequest) DeepCopyInto(out *CertificateRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequest.
func (in *CertificateRequest) DeepCopy() *CertificateRequest {
	if in == nil {
		return nil
	}
	out := new(CertificateRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestCondition) DeepCopyInto(out *CertificateRequestCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestCondition.
func (in *CertificateRequestCondition) DeepCopy() *CertificateRequestCondition {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestList) DeepCopyInto(out *CertificateRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestList.
func (in *CertificateRequestList) DeepCopy() *CertificateRequestList {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestSpec) DeepCopyInto(out *CertificateRequestSpec) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(metav1.Duration)
		**out = **in
	}
	out.IssuerRef = in.IssuerRef
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = make([]KeyUsage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestSpec.
func (in *CertificateRequestSpec) DeepCopy() *CertificateRequestSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestStatus) DeepCopyInto(out *CertificateRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CertificateRequestCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.FailureTime != nil {
		in, out := &in.FailureTime, &out.FailureTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestStatus.
func (in *CertificateRequestStatus) DeepCopy() *CertificateRequestStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(X509Subject)
		(*in).DeepCopyInto(*out)
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.RenewBefore != nil {
		in, out := &in.RenewBefore, &out.RenewBefore
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.DNSNames != nil {
		in, out := &in.DNSNames, &out.DNSNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.URIs != nil {
		in, out := &in.URIs, &out.URIs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EmailAddresses != nil {
		in, out := &in.EmailAddresses, &out.EmailAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Keystores != nil {
		in, out := &in.Keystores, &out.Keystores
		*out = new(CertificateKeystores)
		(*in).DeepCopyInto(*out)
	}
	out.IssuerRef = in.IssuerRef
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = make([]KeyUsage, len(*in))
		copy(*out, *in)
	}
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(CertificatePrivateKey)
		**out = **in
	}
	if in.EncodeUsagesInRequest != nil {
		in, out := &in.EncodeUsagesInRequest, &out.EncodeUsagesInRequest
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStatus) DeepCopyInto(out *CertificateStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CertificateCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastFailureTime != nil {
		in, out := &in.LastFailureTime, &out.LastFailureTime
		*out = (*in).DeepCopy()
	}
	if in.NotBefore != nil {
		in, out := &in.NotBefore, &out.NotBefore
		*out = (*in).DeepCopy()
	}
	if in.NotAfter != nil {
		in, out := &in.NotAfter, &out.NotAfter
		*out = (*in).DeepCopy()
	}
	if in.RenewalTime != nil {
		in, out := &in.RenewalTime, &out.RenewalTime
		*out = (*in).DeepCopy()
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(int)
		**out = **in
	}
	if in.NextPrivateKeySecretName != nil {
		in, out := &in.NextPrivateKeySecretName, &out.NextPrivateKeySecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStatus.
func (in *CertificateStatus) DeepCopy() *CertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIssuer) DeepCopyInto(out *ClusterIssuer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIssuer.
func (in *ClusterIssuer) DeepCopy() *ClusterIssuer {
	if in == nil {
		return nil
	}
	out := new(ClusterIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterIssuer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIssuerList) DeepCopyInto(out *ClusterIssuerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterIssuer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIssuerList.
func (in *ClusterIssuerList) DeepCopy() *ClusterIssuerList {
	if in == nil {
		return nil
	}
	out := new(ClusterIssuerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterIssuerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Issuer) DeepCopyInto(out *Issuer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Issuer.
func (in *Issuer) DeepCopy() *Issuer {
	if in == nil {
		return nil
	}
	out := new(Issuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Issuer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerCondition) DeepCopyInto(out *IssuerCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerCondition.
func (in *IssuerCondition) DeepCopy() *IssuerCondition {
	if in == nil {
		return nil
	}
	out := new(IssuerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerConfig) DeepCopyInto(out *IssuerConfig) {
	*out = *in
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		*out = new(acmev1.ACMEIssuer)
		(*in).DeepCopyInto(*out)
	}
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = new(CAIssuer)
		(*in).DeepCopyInto(*out)
	}
	if in.Vault != nil {
		in, out := &in.Vault, &out.Vault
		*out = new(VaultIssuer)
		(*in).DeepCopyInto(*out)
	}
	if in.SelfSigned != nil {
		in, out := &in.SelfSigned, &out.SelfSigned
		*out = new(SelfSignedIssuer)
		(*in).DeepCopyInto(*out)
	}
	if in.Venafi != nil {
		in, out := &in.Venafi, &out.Venafi
		*out = new(VenafiIssuer)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerConfig.
func (in *IssuerConfig) DeepCopy() *IssuerConfig {
	if in == nil {
		return nil
	}
	out := new(IssuerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerList) DeepCopyInto(out *IssuerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Issuer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerList.
func (in *IssuerList) DeepCopy() *IssuerList {
	if in == nil {
		return nil
	}
	out := new(IssuerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IssuerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerSpec) DeepCopyInto(out *IssuerSpec) {
	*out = *in
	in.IssuerConfig.DeepCopyInto(&out.IssuerConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerSpec.
func (in *IssuerSpec) DeepCopy() *IssuerSpec {
	if in == nil {
		return nil
	}
	out := new(IssuerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerStatus) DeepCopyInto(out *IssuerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]IssuerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		*out = new(acmev1.ACMEIssuerStatus)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerStatus.
func (in *IssuerStatus) DeepCopy() *IssuerStatus {
	if in == nil {
		return nil
	}
	out := new(IssuerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JKSKeystore) DeepCopyInto(out *JKSKeystore) {
	*out = *in
	out.PasswordSecretRef = in.PasswordSecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JKSKeystore.
func (in *JKSKeystore) DeepCopy() *JKSKeystore {
	if in == nil {
		return nil
	}
	out := new(JKSKeystore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PKCS12Keystore) DeepCopyInto(out *PKCS12Keystore) {
	*out = *in
	out.PasswordSecretRef = in.PasswordSecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PKCS12Keystore.
func (in *PKCS12Keystore) DeepCopy() *PKCS12Keystore {
	if in == nil {
		return nil
	}
	out := new(PKCS12Keystore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSignedIssuer) DeepCopyInto(out *SelfSignedIssuer) {
	*out = *in
	if in.CRLDistributionPoints != nil {
		in, out := &in.CRLDistributionPoints, &out.CRLDistributionPoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSignedIssuer.
func (in *SelfSignedIssuer) DeepCopy() *SelfSignedIssuer {
	if in == nil {
		return nil
	}
	out := new(SelfSignedIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAppRole) DeepCopyInto(out *VaultAppRole) {
	*out = *in
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAppRole.
func (in *VaultAppRole) DeepCopy() *VaultAppRole {
	if in == nil {
		return nil
	}
	out := new(VaultAppRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuth) DeepCopyInto(out *VaultAuth) {
	*out = *in
	if in.TokenSecretRef != nil {
		in, out := &in.TokenSecretRef, &out.TokenSecretRef
		*out = new(apismetav1.SecretKeySelector)
		**out = **in
	}
	if in.AppRole != nil {
		in, out := &in.AppRole, &out.AppRole
		*out = new(VaultAppRole)
		**out = **in
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(VaultKubernetesAuth)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuth.
func (in *VaultAuth) DeepCopy() *VaultAuth {
	if in == nil {
		return nil
	}
	out := new(VaultAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultIssuer) DeepCopyInto(out *VaultIssuer) {
	*out = *in
	in.Auth.DeepCopyInto(&out.Auth)
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultIssuer.
func (in *VaultIssuer) DeepCopy() *VaultIssuer {
	if in == nil {
		return nil
	}
	out := new(VaultIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultKubernetesAuth) DeepCopyInto(out *VaultKubernetesAuth) {
	*out = *in
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultKubernetesAuth.
func (in *VaultKubernetesAuth) DeepCopy() *VaultKubernetesAuth {
	if in == nil {
		return nil
	}
	out := new(VaultKubernetesAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VenafiCloud) DeepCopyInto(out *VenafiCloud) {
	*out = *in
	out.APITokenSecretRef = in.APITokenSecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VenafiCloud.
func (in *VenafiCloud) DeepCopy() *VenafiCloud {
	if in == nil {
		return nil
	}
	out := new(VenafiCloud)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VenafiIssuer) DeepCopyInto(out *VenafiIssuer) {
	*out = *in
	if in.TPP != nil {
		in, out := &in.TPP, &out.TPP
		*out = new(VenafiTPP)
		(*in).DeepCopyInto(*out)
	}
	if in.Cloud != nil {
		in, out := &in.Cloud, &out.Cloud
		*out = new(VenafiCloud)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VenafiIssuer.
func (in *VenafiIssuer) DeepCopy() *VenafiIssuer {
	if in == nil {
		return nil
	}
	out := new(VenafiIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VenafiTPP) DeepCopyInto(out *VenafiTPP) {
	*out = *in
	out.CredentialsRef = in.CredentialsRef
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VenafiTPP.
func (in *VenafiTPP) DeepCopy() *VenafiTPP {
	if in == nil {
		return nil
	}
	out := new(VenafiTPP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *X509Subject) DeepCopyInto(out *X509Subject) {
	*out = *in
	if in.Organizations != nil {
		in, out := &in.Organizations, &out.Organizations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Countries != nil {
		in, out := &in.Countries, &out.Countries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OrganizationalUnits != nil {
		in, out := &in.OrganizationalUnits, &out.OrganizationalUnits
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Localities != nil {
		in, out := &in.Localities, &out.Localities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Provinces != nil {
		in, out := &in.Provinces, &out.Provinces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StreetAddresses != nil {
		in, out := &in.StreetAddresses, &out.StreetAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostalCodes != nil {
		in, out := &in.PostalCodes, &out.PostalCodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new X509Subject.
func (in *X509Subject) DeepCopy() *X509Subject {
	if in == nil {
		return nil
	}
	out := new(X509Subject)
	in.DeepCopyInto(out)
	return out
}
