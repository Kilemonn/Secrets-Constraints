package credential_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentialProviderIdentifier_FromString_InvalidString(t *testing.T) {
	identifier := CredentialProviderIdentifierFromString("Does not exist")
	assert.Equal(t, CredentialProviderIdentifierInvalid, identifier)
}

func TestCredentialProviderIdentifier_IsValid(t *testing.T) {
	values := credentialProviderIdentifierValues()
	for _, val := range values {
		assert.Equal(t, val != CredentialProviderIdentifierInvalid, val.IsValid())
	}
}

func TestCredentialProviderIdentifier(t *testing.T) {
	values := credentialProviderIdentifierValues()
	labels := credentialProviderIdentifierStrings()

	assert.Equal(t, len(values), len(labels))
	for i := range len(values) {
		fromString := CredentialProviderIdentifierFromString(labels[i])
		assert.Equal(t, fromString, values[i])
		assert.Equal(t, uint(i), fromString.Index())
		assert.Equal(t, fromString.String(), labels[i])
	}
}
