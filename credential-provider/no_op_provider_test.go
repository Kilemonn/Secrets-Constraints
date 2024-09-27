package credential_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoOpProvider(t *testing.T) {
	provider := NewNoOpProvider()
	assert.Empty(t, provider.GetCredentials())
	assert.Empty(t, provider.GetCredentialNames())
	assert.Empty(t, provider.GetCredentialWithName("test"))
}
