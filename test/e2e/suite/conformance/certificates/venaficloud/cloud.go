/*
Copyright 2020 The Jetstack cert-manager contributors.

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

package venaficloud

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	cmapi "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	cmmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	"github.com/jetstack/cert-manager/test/e2e/framework"
	"github.com/jetstack/cert-manager/test/e2e/framework/helper/validations"
	"github.com/jetstack/cert-manager/test/e2e/framework/util/errors"
	"github.com/jetstack/cert-manager/test/e2e/suite/conformance/certificates"
	vaddon "github.com/jetstack/cert-manager/test/e2e/suite/issuers/venafi/addon"
)

var _ = framework.ConformanceDescribe("Certificates", func() {
	// unsupportedFeatures is a list of features that are not supported by the
	// Venafi Cloud issuer.
	var unsupportedFeatures = certificates.NewFeatureSet(
		certificates.DurationFeature,
		// Due to the current configuration of the test environment, it does not
		// support signing certificates that pair with an elliptic curve private
		// key or using the same private key multiple times.
		certificates.ECDSAFeature,
		certificates.EmailSANsFeature,
		certificates.CommonNameFeature,
		certificates.IPAddressFeature,
		certificates.URISANsFeature,
	)

	provisioner := new(venafiProvisioner)
	(&certificates.Suite{
		Name:                      "Venafi Cloud Issuer",
		CreateIssuerFunc:          provisioner.createIssuer,
		DeleteIssuerFunc:          provisioner.delete,
		UnsupportedFeatures:       unsupportedFeatures,
		ValidateCertificateChecks: validations.DefaultCertificateValidations,
	}).Define()

	(&certificates.Suite{
		Name:                      "Venafi Cloud ClusterIssuer",
		CreateIssuerFunc:          provisioner.createClusterIssuer,
		DeleteIssuerFunc:          provisioner.delete,
		UnsupportedFeatures:       unsupportedFeatures,
		ValidateCertificateChecks: validations.DefaultCertificateValidations,
	}).Define()
})

type venafiProvisioner struct {
	cloud *vaddon.VenafiCloud
}

func (v *venafiProvisioner) delete(f *framework.Framework, ref cmmeta.ObjectReference) {
	Expect(v.cloud.Deprovision()).NotTo(HaveOccurred(), "failed to deprovision cloud venafi")

	if ref.Kind == "ClusterIssuer" {
		err := f.CertManagerClientSet.CertmanagerV1alpha2().ClusterIssuers().Delete(context.TODO(), ref.Name, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred())
	}
}

func (v *venafiProvisioner) createIssuer(f *framework.Framework) cmmeta.ObjectReference {
	By("Creating a Venafi Cloud Issuer")

	v.cloud = &vaddon.VenafiCloud{
		Namespace: f.Namespace.Name,
	}

	err := v.cloud.Setup(f.Config)
	if errors.IsSkip(err) {
		framework.Skipf("Skipping test as addon could not be setup: %v", err)
	}
	Expect(err).NotTo(HaveOccurred(), "failed to provision venafi cloud issuer")

	Expect(v.cloud.Provision()).NotTo(HaveOccurred(), "failed to provision tpp venafi")

	issuer := v.cloud.Details().BuildIssuer()
	issuer, err = f.CertManagerClientSet.CertmanagerV1alpha2().Issuers(f.Namespace.Name).Create(context.TODO(), issuer, metav1.CreateOptions{})
	Expect(err).NotTo(HaveOccurred(), "failed to create issuer for venafi")

	return cmmeta.ObjectReference{
		Group: cmapi.SchemeGroupVersion.Group,
		Kind:  cmapi.IssuerKind,
		Name:  issuer.Name,
	}
}

func (v *venafiProvisioner) createClusterIssuer(f *framework.Framework) cmmeta.ObjectReference {
	By("Creating a Venafi ClusterIssuer")

	v.cloud = &vaddon.VenafiCloud{
		Namespace: f.Config.Addons.CertManager.ClusterResourceNamespace,
	}

	err := v.cloud.Setup(f.Config)
	if errors.IsSkip(err) {
		framework.Skipf("Skipping test as addon could not be setup: %v", err)
	}
	Expect(err).NotTo(HaveOccurred(), "failed to setup tpp venafi")

	Expect(v.cloud.Provision()).NotTo(HaveOccurred(), "failed to provision tpp venafi")

	issuer := v.cloud.Details().BuildClusterIssuer()
	issuer, err = f.CertManagerClientSet.CertmanagerV1alpha2().ClusterIssuers().Create(context.TODO(), issuer, metav1.CreateOptions{})
	Expect(err).NotTo(HaveOccurred(), "failed to create issuer for venafi")

	return cmmeta.ObjectReference{
		Group: cmapi.SchemeGroupVersion.Group,
		Kind:  cmapi.ClusterIssuerKind,
		Name:  issuer.Name,
	}
}