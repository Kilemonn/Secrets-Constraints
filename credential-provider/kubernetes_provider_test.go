package credential_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKubernetesProvider(t *testing.T) {
	t.Skip("Skipping since we cannot guarantee the local kubernetes environment is correctly setup")
	m := make(map[string]interface{})
	m[property_namespace] = "default"
	m[property_secret_name] = "my-kubernetes-secret"
	provider, err := NewKubernetesProvider(m)
	assert.NoError(t, err)

	val, err := provider.GetCredentialWithName("my-kubernetes-secret")
	assert.NoError(t, err)
	assert.NotEmpty(t, val)
}
