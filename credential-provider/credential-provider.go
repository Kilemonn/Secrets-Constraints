package credential_provider

import (
	"slices"
	"strings"
)

type CredentialProviderIdentifier uint

const (
	CredentialProviderIdentifierInvalid    CredentialProviderIdentifier = iota
	CredentialProviderIdentifierGCP        CredentialProviderIdentifier = iota
	CredentialProviderIdentifierAWS        CredentialProviderIdentifier = iota
	CredentialProviderIdentifierENV        CredentialProviderIdentifier = iota
	CredentialProviderIdentifierKubernetes CredentialProviderIdentifier = iota

	invalid    = "invalid"
	gcp        = "gcp"
	aws        = "aws"
	env        = "env"
	kubernetes = "kubernetes"
)

func credentialProviderIdentifierValues() []CredentialProviderIdentifier {
	return []CredentialProviderIdentifier{
		CredentialProviderIdentifierInvalid,
		CredentialProviderIdentifierGCP,
		CredentialProviderIdentifierAWS,
		CredentialProviderIdentifierENV,
		CredentialProviderIdentifierKubernetes,
	}
}

func credentialProviderIdentifierStrings() []string {
	return []string{invalid, gcp, aws, env, kubernetes}
}

func CredentialProviderIdentifierFromString(input string) CredentialProviderIdentifier {
	index := slices.Index(credentialProviderIdentifierStrings(), strings.ToLower(input))
	if index == -1 {
		return CredentialProviderIdentifierInvalid
	} else {
		return CredentialProviderIdentifier(index)
	}
}

func (c CredentialProviderIdentifier) String() string {
	return credentialProviderIdentifierStrings()[c.Index()]
}

func (c CredentialProviderIdentifier) Index() uint {
	return uint(c)
}

func (c CredentialProviderIdentifier) IsValid() bool {
	return c.Index() != CredentialProviderIdentifierInvalid.Index()
}

func IsValidProvider(providerName string) bool {
	return CredentialProviderIdentifierFromString(providerName).IsValid()
}

type CredentialProvider struct {
	Identifier CredentialProviderIdentifier
}

type Credential struct {
	Name  string
	Value string
}

type Provider interface {
	GetCredentials() []Credential
	GetCredentialNames() []string
	GetCredentialWithName(string) string
}
