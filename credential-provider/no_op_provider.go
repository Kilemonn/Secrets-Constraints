package credential_provider

type NoOpProvider struct {
}

func NewNoOpProvider() (provider NoOpProvider) {
	return
}

func (p NoOpProvider) GetCredentialNames() ([]string, error) {
	return []string{}, nil
}

func (p NoOpProvider) GetCredentialWithName(key string) (string, error) {
	return "", nil
}

func (p NoOpProvider) Shutdown() {
	// No-op
}
