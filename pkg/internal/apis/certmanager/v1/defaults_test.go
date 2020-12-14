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

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	cmapi "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	cmmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
)

func TestSetDefaultsVenafiCloud(t *testing.T) {
	t.Run("set-defaults", func(t *testing.T) {
		o := &cmapi.VenafiCloud{}
		SetDefaults_VenafiCloud(o)
		assert.Equal(t, cmapi.DefaultVenafiCloudURL, o.URL)
		assert.Equal(t, cmapi.DefaultVenafiCloudAPITokenSecretRefKey, o.APITokenSecretRef.Key)
	})

	t.Run("no-clobber", func(t *testing.T) {
		o := &cmapi.VenafiCloud{
			URL: "https://custom-url",
			APITokenSecretRef: cmmeta.SecretKeySelector{
				Key: "foo",
			},
		}
		SetDefaults_VenafiCloud(o)
		assert.Equal(t, "https://custom-url", o.URL)
		assert.Equal(t, "foo", o.APITokenSecretRef.Key)
	})
}
