package v1alpha1

import (
    "k8s.io/apimachinery/pkg/runtime/schema"
    "k8s.io/client-go/kubernetes/scheme"
    "k8s.io/client-go/rest"
)

type ExampleV1Alpha1Interface interface {
    RESTClient() rest.Interface
    Migration(namespace string) MigrationInterface
    MigrationPod(namespace string) MigrationPodInterface
}

type ExampleV1Alpha1Client struct {
    restClient rest.Interface
    migration  *MigrationClient
    pod        MigrationPodClient
}

func NewForConfig(c *rest.Config) (*ExampleV1Alpha1Client, error) {
    config := *c
    config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "keti.migration", Version: "v1alpha1"}
    config.APIPath = "/apis"
    config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
    config.UserAgent = rest.DefaultKubernetesUserAgent()

    client, err := rest.RESTClientFor(&config)
    if err != nil {
        return nil, err
    }

    return &ExampleV1Alpha1Client{restClient: client}, nil
}

func (c *ExampleV1Alpha1Client) RESTClient() rest.Interface {
    if c == nil {
        return nil
    }
    return c.restClient
}

func (c *ExampleV1Alpha1Client) Migration(namespace string) MigrationInterface {
    return newMigrationClient(c,namespace)
}

func (c *ExampleV1Alpha1Client) MigrationPod(namespace string) MigrationPodInterface {
    return newMigrationPodClient(c,namespace)
}