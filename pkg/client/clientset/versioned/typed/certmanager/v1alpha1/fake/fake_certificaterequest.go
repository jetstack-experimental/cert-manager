/*
Copyright 2019 The Jetstack cert-manager contributors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCertificateRequests implements CertificateRequestInterface
type FakeCertificateRequests struct {
	Fake *FakeCertmanagerV1alpha1
}

var certificaterequestsResource = schema.GroupVersionResource{Group: "certmanager.k8s.io", Version: "v1alpha1", Resource: "certificaterequests"}

var certificaterequestsKind = schema.GroupVersionKind{Group: "certmanager.k8s.io", Version: "v1alpha1", Kind: "CertificateRequest"}

// Get takes name of the certificateRequest, and returns the corresponding certificateRequest object, and an error if there is any.
func (c *FakeCertificateRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.CertificateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(certificaterequestsResource, name), &v1alpha1.CertificateRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateRequest), err
}

// List takes label and field selectors, and returns the list of CertificateRequests that match those selectors.
func (c *FakeCertificateRequests) List(opts v1.ListOptions) (result *v1alpha1.CertificateRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(certificaterequestsResource, certificaterequestsKind, opts), &v1alpha1.CertificateRequestList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CertificateRequestList{ListMeta: obj.(*v1alpha1.CertificateRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.CertificateRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificateRequests.
func (c *FakeCertificateRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(certificaterequestsResource, opts))
}

// Create takes the representation of a certificateRequest and creates it.  Returns the server's representation of the certificateRequest, and an error, if there is any.
func (c *FakeCertificateRequests) Create(certificateRequest *v1alpha1.CertificateRequest) (result *v1alpha1.CertificateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(certificaterequestsResource, certificateRequest), &v1alpha1.CertificateRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateRequest), err
}

// Update takes the representation of a certificateRequest and updates it. Returns the server's representation of the certificateRequest, and an error, if there is any.
func (c *FakeCertificateRequests) Update(certificateRequest *v1alpha1.CertificateRequest) (result *v1alpha1.CertificateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(certificaterequestsResource, certificateRequest), &v1alpha1.CertificateRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCertificateRequests) UpdateStatus(certificateRequest *v1alpha1.CertificateRequest) (*v1alpha1.CertificateRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(certificaterequestsResource, "status", certificateRequest), &v1alpha1.CertificateRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateRequest), err
}

// Delete takes name of the certificateRequest and deletes it. Returns an error if one occurs.
func (c *FakeCertificateRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(certificaterequestsResource, name), &v1alpha1.CertificateRequest{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificateRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(certificaterequestsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CertificateRequestList{})
	return err
}

// Patch applies the patch and returns the patched certificateRequest.
func (c *FakeCertificateRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CertificateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(certificaterequestsResource, name, pt, data, subresources...), &v1alpha1.CertificateRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateRequest), err
}
