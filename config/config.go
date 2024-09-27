package config

import (
	"fmt"
	"os"

	"github.com/Kilemonn/Secrets-Constraints/constraint"
	credential_provider "github.com/Kilemonn/Secrets-Constraints/credential-provider"
	"gopkg.in/yaml.v3"
)

const (
	yamlPropertyConstraints         = "constraints"
	yamlPropertyCredentialProviders = "credential-providers"

	yamlPropertyCondition = "condition"
	yamlPropertyPattern   = "pattern"
)

func ValidateConfiguration(configFilePath string) (providers []credential_provider.CredentialProvider, constraints []constraint.Constraint, err error) {
	// TODO: Introduce strict/relaxed flag to fail during setup if constraints or providers cannot be resolved

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Failed to read in file [%s]. Error: [%s].\n", configFilePath, err.Error())
		return
	}

	var unmarshalled = make(map[string]interface{})
	err = yaml.Unmarshal(data, &unmarshalled)
	if err != nil {
		fmt.Printf("Failed to unmarshal data from file [%s]. Error: [%s].\n", configFilePath, err.Error())
		return
	}

	if credProviders, exists := unmarshalled[yamlPropertyCredentialProviders]; exists {
		providers, err = getProviders(credProviders.([]interface{}))
		if err != nil {
			fmt.Printf("Failed to validate [%s]. Error: [%s].\n", yamlPropertyCredentialProviders, err.Error())
			return
		}
	} else {
		fmt.Printf("Failed to validate configuration yaml [%s] does not exist.\n", yamlPropertyCredentialProviders)
		return
	}

	if constraintsMap, exists := unmarshalled[yamlPropertyConstraints]; exists {
		constraints, err = getConstraints(constraintsMap.([]interface{}))
		if err != nil {
			fmt.Printf("Failed to validate [%s]. Error: [%s].\n", yamlPropertyConstraints, err.Error())
			return
		}
	} else {
		fmt.Printf("Failed to validate configuration yaml [%s] does not exist.\n", yamlPropertyConstraints)
		return
	}

	return
}

func getProviders(credProviders []interface{}) (credentialProviders []credential_provider.CredentialProvider, err error) {
	for _, val := range credProviders {
		providerId := credential_provider.CredentialProviderIdentifierFromString(val.(string))
		if !providerId.IsValid() {
			fmt.Printf("Failed to register credential provider with name [%s], pleaese provide only valid provider names.\n", val)
		} else {
			provider := credential_provider.CredentialProvider{
				Identifier: providerId,
			}
			credentialProviders = append(credentialProviders, provider)
			fmt.Printf("Registered credential provider with ID [%s].\n", providerId.String())
		}
	}
	return
}

func getConstraints(constraintsMap []interface{}) (constraints []constraint.Constraint, err error) {
	for _, val := range constraintsMap {
		constraintObj, err := validateConstraint(val.(map[string]interface{}))
		if err != nil {
			fmt.Printf("Failed to validate constraint: [%s].\n", err.Error())
		} else {
			constraints = append(constraints, constraintObj)
			fmt.Printf("Registered constraint with name [%s].\n", constraintObj.Name)
		}
	}
	return
}

func validateConstraint(constraintMap map[string]interface{}) (constraintObj constraint.Constraint, err error) {
	if len(constraintMap) != 1 {
		fmt.Println("Expected only len of 1 holding the constraints name.")
		err = fmt.Errorf("found more than one contraint name")
		return
	}

	for name, properties := range constraintMap {
		notContainedKeys := containsAllKeys([]string{yamlPropertyCondition, yamlPropertyPattern}, properties.(map[string]interface{}))
		if len(notContainedKeys) != 0 {
			fmt.Printf("Required properties [%s] are not defined for constraint with name [%s].", notContainedKeys, name)
			err = fmt.Errorf("required properties [%s] are not defined for constraint with name [%s]", notContainedKeys, name)
			return
		}

		propertiesAsMap := properties.(map[string]interface{})
		constraintObj, err = constraint.NewConstraint(name, propertiesAsMap[yamlPropertyPattern].(string), propertiesAsMap[yamlPropertyCondition].(string))
		if err != nil {
			fmt.Printf("Failed to construct constraint [%s] due to [%s].\n", name, err.Error())
		}
	}

	return
}

func containsAllKeys(keys []string, m map[string]interface{}) (notContainedKeys []string) {
	for _, k := range keys {
		if _, exists := m[k]; !exists {
			notContainedKeys = append(notContainedKeys, k)
		}
	}
	return
}
