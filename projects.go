package v1alpha1

import (
	"github.com/hth0919/migrationcontroller/pkg/apis/keti/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"context"
)

type MigrationInterface interface {
	RESTClient() rest.Interface
	List(opts metav1.ListOptions) (*v1alpha1.MigrationList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.Migration, error)
	Create(migration *v1alpha1.Migration) (*v1alpha1.Migration, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type MigrationClient struct {
	restClient rest.Interface
	ns         string
}

func newMigrationClient(c *ExampleV1Alpha1Client, namespace string) *MigrationClient {
	return &MigrationClient{
		restClient: c.RESTClient(),
		ns:     namespace,
	}
}

func (c *MigrationClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *MigrationClient) List(opts metav1.ListOptions) (*v1alpha1.MigrationList, error) {
	result := v1alpha1.MigrationList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("rules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *MigrationClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.Migration, error) {
	result := v1alpha1.Migration{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("rules").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *MigrationClient) Create(migration *v1alpha1.Migration) (*v1alpha1.Migration, error) {
	result := v1alpha1.Migration{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("rules").
		Body(migration).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *MigrationClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("rules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}