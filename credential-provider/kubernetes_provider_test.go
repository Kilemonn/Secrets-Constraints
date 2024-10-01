package credential_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKubernetesProvider(t *testing.T) {
	m := make(map[string]interface{})
	m[property_namespace] = "default"
	m[property_secret_name] = "my-kubernetes-secret"
	provider, err := NewKubernetesProvider(m)
	assert.NoError(t, err)

	val, err := provider.GetCredentialWithName("my-kubernetes-secret")
	assert.NoError(t, err)
	assert.NotEmpty(t, val)
}
