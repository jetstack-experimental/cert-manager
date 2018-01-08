/*
Copyright 2018 Jetstack Ltd.

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
package v1alpha1

import (
	v1alpha1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1"
	scheme "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CertificatesGetter has a method to return a CertificateInterface.
// A group's client should implement this interface.
type CertificatesGetter interface {
	Certificates(namespace string) CertificateInterface
}

// CertificateInterface has methods to work with Certificate resources.
type CertificateInterface interface {
	Create(*v1alpha1.Certificate) (*v1alpha1.Certificate, error)
	Update(*v1alpha1.Certificate) (*v1alpha1.Certificate, error)
	UpdateStatus(*v1alpha1.Certificate) (*v1alpha1.Certificate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Certificate, error)
	List(opts v1.ListOptions) (*v1alpha1.CertificateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Certificate, err error)
	CertificateExpansion
}

// certificates implements CertificateInterface
type certificates struct {
	client rest.Interface
	ns     string
}

// newCertificates returns a Certificates
func newCertificates(c *CertmanagerV1alpha1Client, namespace string) *certificates {
	return &certificates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the certificate, and returns the corresponding certificate object, and an error if there is any.
func (c *certificates) Get(name string, options v1.GetOptions) (result *v1alpha1.Certificate, err error) {
	result = &v1alpha1.Certificate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Certificates that match those selectors.
func (c *certificates) List(opts v1.ListOptions) (result *v1alpha1.CertificateList, err error) {
	result = &v1alpha1.CertificateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested certificates.
func (c *certificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("certificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a certificate and creates it.  Returns the server's representation of the certificate, and an error, if there is any.
func (c *certificates) Create(certificate *v1alpha1.Certificate) (result *v1alpha1.Certificate, err error) {
	result = &v1alpha1.Certificate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("certificates").
		Body(certificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a certificate and updates it. Returns the server's representation of the certificate, and an error, if there is any.
func (c *certificates) Update(certificate *v1alpha1.Certificate) (result *v1alpha1.Certificate, err error) {
	result = &v1alpha1.Certificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("certificates").
		Name(certificate.Name).
		Body(certificate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *certificates) UpdateStatus(certificate *v1alpha1.Certificate) (result *v1alpha1.Certificate, err error) {
	result = &v1alpha1.Certificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("certificates").
		Name(certificate.Name).
		SubResource("status").
		Body(certificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the certificate and deletes it. Returns an error if one occurs.
func (c *certificates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *certificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched certificate.
func (c *certificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Certificate, err error) {
	result = &v1alpha1.Certificate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("certificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
