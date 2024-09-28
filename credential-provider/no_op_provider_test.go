package credential_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoOpProvider(t *testing.T) {
	provider := NewNoOpProvider()
	names, err := provider.GetCredentialNames()
	assert.NoError(t, err)
	assert.Empty(t, names)
	cred, err := provider.GetCredentialWithName("test")
	assert.NoError(t, err)
	assert.Empty(t, cred)
}
