// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	operator_v1 "github.com/openshift/api/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIngressControllers implements IngressControllerInterface
type FakeIngressControllers struct {
	Fake *FakeOperatorV1
	ns   string
}

var ingresscontrollersResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1", Resource: "ingresscontrollers"}

var ingresscontrollersKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1", Kind: "IngressController"}

// Get takes name of the ingressController, and returns the corresponding ingressController object, and an error if there is any.
func (c *FakeIngressControllers) Get(name string, options v1.GetOptions) (result *operator_v1.IngressController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ingresscontrollersResource, c.ns, name), &operator_v1.IngressController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.IngressController), err
}

// List takes label and field selectors, and returns the list of IngressControllers that match those selectors.
func (c *FakeIngressControllers) List(opts v1.ListOptions) (result *operator_v1.IngressControllerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ingresscontrollersResource, ingresscontrollersKind, c.ns, opts), &operator_v1.IngressControllerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operator_v1.IngressControllerList{ListMeta: obj.(*operator_v1.IngressControllerList).ListMeta}
	for _, item := range obj.(*operator_v1.IngressControllerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingressControllers.
func (c *FakeIngressControllers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ingresscontrollersResource, c.ns, opts))

}

// Create takes the representation of a ingressController and creates it.  Returns the server's representation of the ingressController, and an error, if there is any.
func (c *FakeIngressControllers) Create(ingressController *operator_v1.IngressController) (result *operator_v1.IngressController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ingresscontrollersResource, c.ns, ingressController), &operator_v1.IngressController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.IngressController), err
}

// Update takes the representation of a ingressController and updates it. Returns the server's representation of the ingressController, and an error, if there is any.
func (c *FakeIngressControllers) Update(ingressController *operator_v1.IngressController) (result *operator_v1.IngressController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ingresscontrollersResource, c.ns, ingressController), &operator_v1.IngressController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.IngressController), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIngressControllers) UpdateStatus(ingressController *operator_v1.IngressController) (*operator_v1.IngressController, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ingresscontrollersResource, "status", c.ns, ingressController), &operator_v1.IngressController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.IngressController), err
}

// Delete takes name of the ingressController and deletes it. Returns an error if one occurs.
func (c *FakeIngressControllers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ingresscontrollersResource, c.ns, name), &operator_v1.IngressController{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngressControllers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ingresscontrollersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &operator_v1.IngressControllerList{})
	return err
}

// Patch applies the patch and returns the patched ingressController.
func (c *FakeIngressControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *operator_v1.IngressController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ingresscontrollersResource, c.ns, name, data, subresources...), &operator_v1.IngressController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.IngressController), err
}
