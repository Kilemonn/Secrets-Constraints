package credential_provider

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/maps"
)

type EnvironmentProvider struct {
	credentials map[string]string
}

func NewEnvironmentProvider() EnvironmentProvider {
	provider := EnvironmentProvider{
		credentials: make(map[string]string),
	}
	provider.initialiseProvider()

	return provider
}

func (p EnvironmentProvider) initialiseProvider() {
	for _, v := range os.Environ() {
		index := strings.Index(v, "=")
		if index != -1 {
			p.credentials[v[0:index]] = v[index+1:]
		} else {
			fmt.Printf("No '=' found in environment value: [%s]", v)
		}
	}
}

func (p EnvironmentProvider) GetCredentials() map[string]string {
	return p.credentials
}

func (p EnvironmentProvider) GetCredentialNames() []string {
	return maps.Keys(p.credentials)
}

func (p EnvironmentProvider) GetCredentialWithName(key string) string {
	return p.credentials[key]
}
