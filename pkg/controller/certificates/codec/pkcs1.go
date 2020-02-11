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

package codec

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	corev1 "k8s.io/api/core/v1"

	cmmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	"github.com/jetstack/cert-manager/pkg/util/errors"
	"github.com/jetstack/cert-manager/pkg/util/pki"
)

// PKCS1 knows how to encode and decode RSA PKCS.1 formatted private key
// certificate data.
type PKCS1 struct{}

var _ Codec = PKCS1{}

// Encode encodes the given private key into PEM-encoded PKCS.1 format.
// The certificate be encoded into DER format and stored as a PEM,
func (p PKCS1) Encode(d Bundle) (*RawData, error) {
	data := map[string][]byte{}
	rsaPK, ok := d.PrivateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key is not an RSA key")
	}
	pkDER := x509.MarshalPKCS1PrivateKey(rsaPK)
	pkPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: pkDER,
	})
	if pkPEM == nil {
		return nil, fmt.Errorf("failed to encode RSA private key to PEM format")
	}
	data[corev1.TLSPrivateKeyKey] = pkPEM

	if len(d.Certificates) > 0 {
		certsPEM, err := encodeCertificatesASN1PEM(d.Certificates)
		if err != nil {
			return nil, err
		}
		data[corev1.TLSCertKey] = certsPEM
	}

	if len(d.CA) > 0 {
		caPEM, err := encodeCertificatesASN1PEM(d.CA)
		if err != nil {
			return nil, err
		}
		data[cmmeta.TLSCAKey] = caPEM
	}

	return &RawData{Data: data}, nil
}

func (p PKCS1) Decode(e RawData) (*Bundle, error) {
	d := &Bundle{}
	pkPEM := e.Data[corev1.TLSPrivateKeyKey]
	certsPEM := e.Data[corev1.TLSCertKey]
	caPEM := e.Data[cmmeta.TLSCAKey]
	var err error
	if len(pkPEM) > 0 {
		// decode the private key pem
		block, _ := pem.Decode(pkPEM)
		if block == nil {
			return d, errors.NewInvalidData("failed to decode PEM block")
		}
		// "RSA PRIVATE KEY" is the PEM marker used for PKCS1 encoded data
		if block.Type != "RSA PRIVATE KEY" {
			return d, errors.NewInvalidData("unexpected PEM block type %q - PKCS1 data should specify the type as 'RSA PRIVATE KEY'", block.Type)
		}
		pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return d, errors.NewInvalidData("failed to decode private key: %w", err)
		}
		d.PrivateKey = pk
	}
	if len(certsPEM) > 0 {
		d.Certificates, err = pki.DecodeX509CertificateChainBytes(certsPEM)
		if err != nil {
			return d, errors.NewInvalidData(err.Error())
		}
	}
	if len(caPEM) > 0 {
		d.CA, err = pki.DecodeX509CertificateChainBytes(caPEM)
		if err != nil {
			return d, errors.NewInvalidData(err.Error())
		}
	}
	return d, nil
}
