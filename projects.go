package v1alpha1

import (
	"github.com/hth0919/migrationcontroller/pkg/apis/keti/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type MigrationInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.MigrationList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.Migration, error)
	Create(migration *v1alpha1.Migration) (*v1alpha1.Migration, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type migrationClient struct {
	restClient rest.Interface
	ns         string
}

func (c *migrationClient) List(opts metav1.ListOptions) (*v1alpha1.MigrationList, error) {
	result := v1alpha1.MigrationList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("migrations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *migrationClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.Migration, error) {
	result := v1alpha1.Migration{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("migrations").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *migrationClient) Create(migration *v1alpha1.Migration) (*v1alpha1.Migration, error) {
	result := v1alpha1.Migration{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("migrations").
		Body(migration).
		Do().
		Into(&result)

	return &result, err
}

func (c *migrationClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("migrations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}