/*
Copyright 2020 The cert-manager Authors.

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

package metrics

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	cmapi "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	cmmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	logtesting "github.com/jetstack/cert-manager/pkg/logs/testing"
	"github.com/jetstack/cert-manager/test/unit/gen"
)

const expiryMetadata = `
	# HELP certmanager_certificate_expiration_timestamp_seconds The date after which the certificate expires. Expressed as a Unix Epoch Time.
	# TYPE certmanager_certificate_expiration_timestamp_seconds gauge
`

const expiryRelativeMetadata = `
	# HELP certmanager_certificate_expiration_timestamp_seconds_relative The relative time after which the certificate expires. Expressed in seconds.
	# TYPE certmanager_certificate_expiration_timestamp_seconds_relative gauge
`

const readyMetadata = `
  # HELP certmanager_certificate_ready_status The ready status of the certificate.
  # TYPE certmanager_certificate_ready_status gauge
`

func TestCertificateMetrics(t *testing.T) {
	type testT struct {
		crt                                                   *cmapi.Certificate
		timeNow                                               time.Time
		expectedExpiry, expectedRelativeExpiry, expectedReady string
	}
	tests := map[string]testT{
		"certificate with expiry and ready status": {
			timeNow: time.Unix(2208988804-100, 0),
			crt: gen.Certificate("test-certificate",
				gen.SetCertificateNamespace("test-ns"),
				gen.SetCertificateNotAfter(metav1.Time{
					Time: time.Unix(2208988804, 0),
				}),
				gen.SetCertificateStatusCondition(cmapi.CertificateCondition{
					Type:   cmapi.CertificateConditionReady,
					Status: cmmeta.ConditionTrue,
				}),
			),
			expectedExpiry: `
	certmanager_certificate_expiration_timestamp_seconds{name="test-certificate",namespace="test-ns"} 2.208988804e+09
`,
			expectedRelativeExpiry: `
	certmanager_certificate_expiration_timestamp_seconds_relative{name="test-certificate",namespace="test-ns"} 100
`,
			expectedReady: `
        certmanager_certificate_ready_status{condition="False",name="test-certificate",namespace="test-ns"} 0
        certmanager_certificate_ready_status{condition="True",name="test-certificate",namespace="test-ns"} 1
        certmanager_certificate_ready_status{condition="Unknown",name="test-certificate",namespace="test-ns"} 0
`,
		},

		"certificate with no expiry and no status should give an expiry of 0 and Unknown status": {
			crt: gen.Certificate("test-certificate",
				gen.SetCertificateNamespace("test-ns"),
			),
			expectedExpiry: `
	certmanager_certificate_expiration_timestamp_seconds{name="test-certificate",namespace="test-ns"} 0
`,
			expectedRelativeExpiry: `
	certmanager_certificate_expiration_timestamp_seconds_relative{name="test-certificate",namespace="test-ns"} 0
`,
			expectedReady: `
        certmanager_certificate_ready_status{condition="False",name="test-certificate",namespace="test-ns"} 0
        certmanager_certificate_ready_status{condition="True",name="test-certificate",namespace="test-ns"} 0
        certmanager_certificate_ready_status{condition="Unknown",name="test-certificate",namespace="test-ns"} 1
`,
		},

		"certificate with expiry and status False should give an expiry and False status": {
			timeNow: time.Unix(0, 0),
			crt: gen.Certificate("test-certificate",
				gen.SetCertificateNamespace("test-ns"),
				gen.SetCertificateNotAfter(metav1.Time{
					Time: time.Unix(100, 0),
				}),
				gen.SetCertificateStatusCondition(cmapi.CertificateCondition{
					Type:   cmapi.CertificateConditionReady,
					Status: cmmeta.ConditionFalse,
				}),
			),
			expectedExpiry: `
	certmanager_certificate_expiration_timestamp_seconds{name="test-certificate",namespace="test-ns"} 100
`,
			expectedRelativeExpiry: `
	certmanager_certificate_expiration_timestamp_seconds_relative{name="test-certificate",namespace="test-ns"} 100
`,
			expectedReady: `
        certmanager_certificate_ready_status{condition="False",name="test-certificate",namespace="test-ns"} 1
        certmanager_certificate_ready_status{condition="True",name="test-certificate",namespace="test-ns"} 0
        certmanager_certificate_ready_status{condition="Unknown",name="test-certificate",namespace="test-ns"} 0
`,
		},
		"certificate with expiry and status Unknown should give an expiry and Unknown status": {
			timeNow: time.Unix(100000, 0),
			crt: gen.Certificate("test-certificate",
				gen.SetCertificateNamespace("test-ns"),
				gen.SetCertificateNotAfter(metav1.Time{
					Time: time.Unix(99999, 0),
				}),
				gen.SetCertificateStatusCondition(cmapi.CertificateCondition{
					Type:   cmapi.CertificateConditionReady,
					Status: cmmeta.ConditionUnknown,
				}),
			),
			expectedExpiry: `
	certmanager_certificate_expiration_timestamp_seconds{name="test-certificate",namespace="test-ns"} 99999
`,
			expectedRelativeExpiry: `
	certmanager_certificate_expiration_timestamp_seconds_relative{name="test-certificate",namespace="test-ns"} -1
`,
			expectedReady: `
        certmanager_certificate_ready_status{condition="False",name="test-certificate",namespace="test-ns"} 0
        certmanager_certificate_ready_status{condition="True",name="test-certificate",namespace="test-ns"} 0
        certmanager_certificate_ready_status{condition="Unknown",name="test-certificate",namespace="test-ns"} 1
`,
		},
	}
	for n, test := range tests {
		t.Run(n, func(t *testing.T) {
			m := New(logtesting.TestLogger{T: t})
			m.UpdateCertificate(context.TODO(), test.crt, test.timeNow)

			if err := testutil.CollectAndCompare(m.certificateExpiryTimeSeconds,
				strings.NewReader(expiryMetadata+test.expectedExpiry),
				"certmanager_certificate_expiration_timestamp_seconds",
			); err != nil {
				t.Errorf("unexpected collecting result:\n%s", err)
			}

			if err := testutil.CollectAndCompare(m.certificateExpiryTimeSecondsRelative,
				strings.NewReader(expiryRelativeMetadata+test.expectedRelativeExpiry),
				"certmanager_certificate_expiration_timestamp_seconds_relative",
			); err != nil {
				t.Errorf("unexpected collecting result:\n%s", err)
			}

			if err := testutil.CollectAndCompare(m.certificateReadyStatus,
				strings.NewReader(readyMetadata+test.expectedReady),
				"certmanager_certificate_ready_status",
			); err != nil {
				t.Errorf("unexpected collecting result:\n%s", err)
			}
		})
	}
}

func TestCertificateCache(t *testing.T) {
	m := New(logtesting.TestLogger{T: t})

	crt1 := gen.Certificate("crt1",
		gen.SetCertificateUID("uid-1"),
		gen.SetCertificateNotAfter(metav1.Time{
			Time: time.Unix(100, 0),
		}),
		gen.SetCertificateStatusCondition(cmapi.CertificateCondition{
			Type:   cmapi.CertificateConditionReady,
			Status: cmmeta.ConditionUnknown,
		}),
	)
	crt2 := gen.Certificate("crt2",
		gen.SetCertificateUID("uid-2"),
		gen.SetCertificateNotAfter(metav1.Time{
			Time: time.Unix(200, 0),
		}),
		gen.SetCertificateStatusCondition(cmapi.CertificateCondition{
			Type:   cmapi.CertificateConditionReady,
			Status: cmmeta.ConditionTrue,
		}),
	)
	crt3 := gen.Certificate("crt3",
		gen.SetCertificateUID("uid-3"),
		gen.SetCertificateNotAfter(metav1.Time{
			Time: time.Unix(300, 0),
		}),
		gen.SetCertificateStatusCondition(cmapi.CertificateCondition{
			Type:   cmapi.CertificateConditionReady,
			Status: cmmeta.ConditionFalse,
		}),
	)

	// Observe all three Certificate metrics
	timeNow := time.Unix(50, 0)
	m.UpdateCertificate(context.TODO(), crt1, timeNow)
	m.UpdateCertificate(context.TODO(), crt2, timeNow)
	m.UpdateCertificate(context.TODO(), crt3, timeNow)

	// Check all three metrics exist
	if err := testutil.CollectAndCompare(m.certificateReadyStatus,
		strings.NewReader(readyMetadata+`
        certmanager_certificate_ready_status{condition="False",name="crt1",namespace="default-unit-test-ns"} 0
        certmanager_certificate_ready_status{condition="False",name="crt2",namespace="default-unit-test-ns"} 0
        certmanager_certificate_ready_status{condition="False",name="crt3",namespace="default-unit-test-ns"} 1
        certmanager_certificate_ready_status{condition="True",name="crt1",namespace="default-unit-test-ns"} 0
        certmanager_certificate_ready_status{condition="True",name="crt2",namespace="default-unit-test-ns"} 1
        certmanager_certificate_ready_status{condition="True",name="crt3",namespace="default-unit-test-ns"} 0
        certmanager_certificate_ready_status{condition="Unknown",name="crt1",namespace="default-unit-test-ns"} 1
        certmanager_certificate_ready_status{condition="Unknown",name="crt2",namespace="default-unit-test-ns"} 0
        certmanager_certificate_ready_status{condition="Unknown",name="crt3",namespace="default-unit-test-ns"} 0
`),
		"certmanager_certificate_ready_status",
	); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
	if err := testutil.CollectAndCompare(m.certificateExpiryTimeSeconds,
		strings.NewReader(expiryMetadata+`
        certmanager_certificate_expiration_timestamp_seconds{name="crt1",namespace="default-unit-test-ns"} 100
        certmanager_certificate_expiration_timestamp_seconds{name="crt2",namespace="default-unit-test-ns"} 200
        certmanager_certificate_expiration_timestamp_seconds{name="crt3",namespace="default-unit-test-ns"} 300
`),
		"certmanager_certificate_expiration_timestamp_seconds",
	); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
	if err := testutil.CollectAndCompare(m.certificateExpiryTimeSecondsRelative,
		strings.NewReader(expiryRelativeMetadata+`
				certmanager_certificate_expiration_timestamp_seconds_relative{name="crt1",namespace="default-unit-test-ns"} 50
				certmanager_certificate_expiration_timestamp_seconds_relative{name="crt2",namespace="default-unit-test-ns"} 150
				certmanager_certificate_expiration_timestamp_seconds_relative{name="crt3",namespace="default-unit-test-ns"} 250
`),
		"certmanager_certificate_expiration_timestamp_seconds_relative",
	); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}

	// Remove second certificate and check not exists
	m.RemoveCertificate("default-unit-test-ns/crt2")
	if err := testutil.CollectAndCompare(m.certificateReadyStatus,
		strings.NewReader(readyMetadata+`
        certmanager_certificate_ready_status{condition="False",name="crt1",namespace="default-unit-test-ns"} 0
        certmanager_certificate_ready_status{condition="False",name="crt3",namespace="default-unit-test-ns"} 1
        certmanager_certificate_ready_status{condition="True",name="crt1",namespace="default-unit-test-ns"} 0
        certmanager_certificate_ready_status{condition="True",name="crt3",namespace="default-unit-test-ns"} 0
        certmanager_certificate_ready_status{condition="Unknown",name="crt1",namespace="default-unit-test-ns"} 1
        certmanager_certificate_ready_status{condition="Unknown",name="crt3",namespace="default-unit-test-ns"} 0
`),
		"certmanager_certificate_ready_status",
	); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
	if err := testutil.CollectAndCompare(m.certificateExpiryTimeSeconds,
		strings.NewReader(expiryMetadata+`
        certmanager_certificate_expiration_timestamp_seconds{name="crt1",namespace="default-unit-test-ns"} 100
        certmanager_certificate_expiration_timestamp_seconds{name="crt3",namespace="default-unit-test-ns"} 300
`),
		"certmanager_certificate_expiration_timestamp_seconds",
	); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
	if err := testutil.CollectAndCompare(m.certificateExpiryTimeSecondsRelative,
		strings.NewReader(expiryRelativeMetadata+`
				certmanager_certificate_expiration_timestamp_seconds_relative{name="crt1",namespace="default-unit-test-ns"} 50
				certmanager_certificate_expiration_timestamp_seconds_relative{name="crt3",namespace="default-unit-test-ns"} 250
`),
		"certmanager_certificate_expiration_timestamp_seconds_relative",
	); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}

	// Remove all Certificates (even is already removed) and observe no Certificates
	m.RemoveCertificate("default-unit-test-ns/crt1")
	m.RemoveCertificate("default-unit-test-ns/crt2")
	m.RemoveCertificate("default-unit-test-ns/crt3")
	if err := testutil.CollectAndCompare(m.certificateReadyStatus,
		strings.NewReader(readyMetadata),
		"certmanager_certificate_ready_status",
	); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
	if err := testutil.CollectAndCompare(m.certificateExpiryTimeSeconds,
		strings.NewReader(expiryMetadata),
		"certmanager_certificate_expiration_timestamp_seconds",
	); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
}
