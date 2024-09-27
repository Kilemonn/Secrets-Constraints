package credential_provider

type NoOpProvider struct {
}

func NewNoOpProvider() (provider NoOpProvider) {
	provider.initialiseProvider()
	return provider
}

func (p NoOpProvider) initialiseProvider() {
	// No-op
}

func (p NoOpProvider) GetCredentials() (creds map[string]string) {
	return creds
}

func (p NoOpProvider) GetCredentialNames() (names []string) {
	return names
}

func (p NoOpProvider) GetCredentialWithName(key string) string {
	return ""
}
