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

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/conversion"

	"github.com/jetstack/cert-manager/pkg/apis/acme/v1alpha3"
	"github.com/jetstack/cert-manager/pkg/internal/apis/acme"
)

func Convert_v1alpha3_ChallengeSpec_To_acme_ChallengeSpec(in *v1alpha3.ChallengeSpec, out *acme.ChallengeSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha3_ChallengeSpec_To_acme_ChallengeSpec(in, out, s); err != nil {
		return err
	}

	out.AuthorizationURL = in.AuthzURL

	return nil
}

func Convert_acme_ChallengeSpec_To_v1alpha3_ChallengeSpec(in *acme.ChallengeSpec, out *v1alpha3.ChallengeSpec, s conversion.Scope) error {
	if err := autoConvert_acme_ChallengeSpec_To_v1alpha3_ChallengeSpec(in, out, s); err != nil {
		return err
	}

	out.AuthzURL = in.AuthorizationURL

	return nil
}

func Convert_v1alpha3_OrderSpec_To_acme_OrderSpec(in *v1alpha3.OrderSpec, out *acme.OrderSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha3_OrderSpec_To_acme_OrderSpec(in, out, s); err != nil {
		return err
	}

	out.Request = in.CSR

	return nil
}

func Convert_acme_OrderSpec_To_v1alpha3_OrderSpec(in *acme.OrderSpec, out *v1alpha3.OrderSpec, s conversion.Scope) error {
	if err := autoConvert_acme_OrderSpec_To_v1alpha3_OrderSpec(in, out, s); err != nil {
		return err
	}

	out.CSR = in.Request

	return nil
}