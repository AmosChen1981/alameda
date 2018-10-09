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
	v1alpha1 "github.com/containers-ai/alameda/pkg/apis/autoscaling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlamedaVPAs implements AlamedaVPAInterface
type FakeAlamedaVPAs struct {
	Fake *FakeAutoscalingV1alpha1
	ns   string
}

var alamedavpasResource = schema.GroupVersionResource{Group: "autoscaling.containers.ai", Version: "v1alpha1", Resource: "alamedavpas"}

var alamedavpasKind = schema.GroupVersionKind{Group: "autoscaling.containers.ai", Version: "v1alpha1", Kind: "AlamedaVPA"}

// Get takes name of the alamedaVPA, and returns the corresponding alamedaVPA object, and an error if there is any.
func (c *FakeAlamedaVPAs) Get(name string, options v1.GetOptions) (result *v1alpha1.AlamedaVPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alamedavpasResource, c.ns, name), &v1alpha1.AlamedaVPA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlamedaVPA), err
}

// List takes label and field selectors, and returns the list of AlamedaVPAs that match those selectors.
func (c *FakeAlamedaVPAs) List(opts v1.ListOptions) (result *v1alpha1.AlamedaVPAList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alamedavpasResource, alamedavpasKind, c.ns, opts), &v1alpha1.AlamedaVPAList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AlamedaVPAList{ListMeta: obj.(*v1alpha1.AlamedaVPAList).ListMeta}
	for _, item := range obj.(*v1alpha1.AlamedaVPAList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alamedaVPAs.
func (c *FakeAlamedaVPAs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alamedavpasResource, c.ns, opts))

}

// Create takes the representation of a alamedaVPA and creates it.  Returns the server's representation of the alamedaVPA, and an error, if there is any.
func (c *FakeAlamedaVPAs) Create(alamedaVPA *v1alpha1.AlamedaVPA) (result *v1alpha1.AlamedaVPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alamedavpasResource, c.ns, alamedaVPA), &v1alpha1.AlamedaVPA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlamedaVPA), err
}

// Update takes the representation of a alamedaVPA and updates it. Returns the server's representation of the alamedaVPA, and an error, if there is any.
func (c *FakeAlamedaVPAs) Update(alamedaVPA *v1alpha1.AlamedaVPA) (result *v1alpha1.AlamedaVPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alamedavpasResource, c.ns, alamedaVPA), &v1alpha1.AlamedaVPA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlamedaVPA), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlamedaVPAs) UpdateStatus(alamedaVPA *v1alpha1.AlamedaVPA) (*v1alpha1.AlamedaVPA, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alamedavpasResource, "status", c.ns, alamedaVPA), &v1alpha1.AlamedaVPA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlamedaVPA), err
}

// Delete takes name of the alamedaVPA and deletes it. Returns an error if one occurs.
func (c *FakeAlamedaVPAs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(alamedavpasResource, c.ns, name), &v1alpha1.AlamedaVPA{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlamedaVPAs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alamedavpasResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AlamedaVPAList{})
	return err
}

// Patch applies the patch and returns the patched alamedaVPA.
func (c *FakeAlamedaVPAs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AlamedaVPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alamedavpasResource, c.ns, name, data, subresources...), &v1alpha1.AlamedaVPA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlamedaVPA), err
}
