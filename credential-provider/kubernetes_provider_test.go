package credential_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKubernetesProvider(t *testing.T) {
	provider, err := NewKubernetesProvider()
	assert.NoError(t, err)

	val, err := provider.GetCredentialWithName("my-kubernetes-secret")
	assert.NoError(t, err)
	assert.NotEmpty(t, val)
}
