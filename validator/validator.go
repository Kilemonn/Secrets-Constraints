package validator

import (
	"github.com/Kilemonn/Secrets-Constraints/constraint"
	credential_provider "github.com/Kilemonn/Secrets-Constraints/credential-provider"
)

func ExecuteConstraintsAgainstProviders(providers []credential_provider.CredentialProvider, constraints []constraint.Constraint) (failed map[string][]string) {

	return
}
