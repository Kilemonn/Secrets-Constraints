package credential_provider

import (
	"context"
	"fmt"
	"path/filepath"

	"golang.org/x/exp/maps"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type KubernetesProvider struct {
	ctx       context.Context
	client    kubernetes.Interface
	namespace string
}

func NewKubernetesProvider() (provider KubernetesProvider, err error) {
	if home := homedir.HomeDir(); home != "" {
		kubeConfigPath := filepath.Join(home, ".kube", "config")
		var config *rest.Config
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfigPath)
		if err != nil {
			return
		}

		provider.client, err = kubernetes.NewForConfig(config)
		if err != nil {
			return
		}
		provider.namespace = "default"
	} else {
		err = fmt.Errorf("unable to find .kube/config file")
	}

	return
}

func (p KubernetesProvider) GetCredentialNames() ([]string, error) {
	names := []string{}
	client := p.client.CoreV1()
	secrets := client.Secrets(p.namespace)
	list, err := secrets.List(p.ctx, v1.ListOptions{})
	if err != nil {
		return names, err
	}

	for _, item := range list.Items {
		names = append(names, maps.Values(item.StringData)[0])
	}
	return names, nil
}

func (p KubernetesProvider) GetCredentialWithName(key string) (string, error) {
	client := p.client.CoreV1()
	secrets := client.Secrets(p.namespace)
	secret, err := secrets.Get(p.ctx, key, v1.GetOptions{})
	if err != nil {
		return "", err
	}

	return maps.Values(secret.StringData)[0], nil
}

func (p KubernetesProvider) Shutdown() {
	// No-op
}
