/*
Copyright The Kubernetes Authors.

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
	context "context"
	json "encoding/json"
	fmt "fmt"

	v1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	admissionregistrationv1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1"
	testing "k8s.io/client-go/testing"
)

// FakeValidatingWebhookConfigurations implements ValidatingWebhookConfigurationInterface
type FakeValidatingWebhookConfigurations struct {
	Fake *FakeAdmissionregistrationV1
}

var validatingwebhookconfigurationsResource = v1.SchemeGroupVersion.WithResource("validatingwebhookconfigurations")

var validatingwebhookconfigurationsKind = v1.SchemeGroupVersion.WithKind("ValidatingWebhookConfiguration")

// Get takes name of the validatingWebhookConfiguration, and returns the corresponding validatingWebhookConfiguration object, and an error if there is any.
func (c *FakeValidatingWebhookConfigurations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ValidatingWebhookConfiguration, err error) {
	emptyResult := &v1.ValidatingWebhookConfiguration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(validatingwebhookconfigurationsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ValidatingWebhookConfiguration), err
}

// List takes label and field selectors, and returns the list of ValidatingWebhookConfigurations that match those selectors.
func (c *FakeValidatingWebhookConfigurations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ValidatingWebhookConfigurationList, err error) {
	emptyResult := &v1.ValidatingWebhookConfigurationList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(validatingwebhookconfigurationsResource, validatingwebhookconfigurationsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ValidatingWebhookConfigurationList{ListMeta: obj.(*v1.ValidatingWebhookConfigurationList).ListMeta}
	for _, item := range obj.(*v1.ValidatingWebhookConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested validatingWebhookConfigurations.
func (c *FakeValidatingWebhookConfigurations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(validatingwebhookconfigurationsResource, opts))
}

// Create takes the representation of a validatingWebhookConfiguration and creates it.  Returns the server's representation of the validatingWebhookConfiguration, and an error, if there is any.
func (c *FakeValidatingWebhookConfigurations) Create(ctx context.Context, validatingWebhookConfiguration *v1.ValidatingWebhookConfiguration, opts metav1.CreateOptions) (result *v1.ValidatingWebhookConfiguration, err error) {
	emptyResult := &v1.ValidatingWebhookConfiguration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(validatingwebhookconfigurationsResource, validatingWebhookConfiguration, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ValidatingWebhookConfiguration), err
}

// Update takes the representation of a validatingWebhookConfiguration and updates it. Returns the server's representation of the validatingWebhookConfiguration, and an error, if there is any.
func (c *FakeValidatingWebhookConfigurations) Update(ctx context.Context, validatingWebhookConfiguration *v1.ValidatingWebhookConfiguration, opts metav1.UpdateOptions) (result *v1.ValidatingWebhookConfiguration, err error) {
	emptyResult := &v1.ValidatingWebhookConfiguration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(validatingwebhookconfigurationsResource, validatingWebhookConfiguration, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ValidatingWebhookConfiguration), err
}

// Delete takes name of the validatingWebhookConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeValidatingWebhookConfigurations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(validatingwebhookconfigurationsResource, name, opts), &v1.ValidatingWebhookConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeValidatingWebhookConfigurations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(validatingwebhookconfigurationsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ValidatingWebhookConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched validatingWebhookConfiguration.
func (c *FakeValidatingWebhookConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ValidatingWebhookConfiguration, err error) {
	emptyResult := &v1.ValidatingWebhookConfiguration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(validatingwebhookconfigurationsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ValidatingWebhookConfiguration), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied validatingWebhookConfiguration.
func (c *FakeValidatingWebhookConfigurations) Apply(ctx context.Context, validatingWebhookConfiguration *admissionregistrationv1.ValidatingWebhookConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ValidatingWebhookConfiguration, err error) {
	if validatingWebhookConfiguration == nil {
		return nil, fmt.Errorf("validatingWebhookConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(validatingWebhookConfiguration)
	if err != nil {
		return nil, err
	}
	name := validatingWebhookConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("validatingWebhookConfiguration.Name must be provided to Apply")
	}
	emptyResult := &v1.ValidatingWebhookConfiguration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(validatingwebhookconfigurationsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ValidatingWebhookConfiguration), err
}
