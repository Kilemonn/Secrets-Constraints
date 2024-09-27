package validator

import (
	"github.com/Kilemonn/Secrets-Constraints/constraint"
	credential_provider "github.com/Kilemonn/Secrets-Constraints/credential-provider"
)

func ExecuteConstraintsAgainstProviders(providers []credential_provider.CredentialProvider, constraints []constraint.Constraint) map[string][]string {
	failed := make(map[string][]string)

	for _, provider := range providers {
		for _, credentialName := range provider.Provider.GetCredentialNames() {
			for _, constraint := range constraints {
				if constraint.Pattern.Matches(credentialName) {
					if !constraint.Condition.ApplyCondition(provider.Provider.GetCredentialWithName(credentialName)) {
						// fmt.Printf("Fail - Provider [%s], Constraint [%s], Credential [%s].\n", provider.Identifier.String(), constraint.Name, credentialName)
						if _, exists := failed[constraint.Name]; !exists {
							failed[constraint.Name] = make([]string, 0)
						}
						failed[constraint.Name] = append(failed[constraint.Name], credentialName)
					} else {
						// fmt.Printf("Pass - Provider [%s], Constraint [%s], Credential [%s].\n", provider.Identifier.String(), constraint.Name, credentialName)
					}
				}
			}
		}
	}
	return failed
}
