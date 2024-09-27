package main

import (
	"flag"
	"fmt"

	"github.com/Kilemonn/Secrets-Constraints/config"
	"github.com/Kilemonn/Secrets-Constraints/consts"
	"github.com/Kilemonn/Secrets-Constraints/validator"
)

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, consts.ARG_FILE_PATH, "", "configuration file path")
	flag.Parse()

	if configFilePath == "" {
		fmt.Printf("Expected value for flag [%s].\n", consts.ARG_FILE_PATH)
		return
	}

	providers, constraints, err := config.ValidateConfiguration(configFilePath)
	if err != nil {
		fmt.Printf("Failed to initialise constraints and/or providers. Error: [%s].\n", err.Error())
	}

	failedConstraints := validator.ExecuteConstraintsAgainstProviders(providers, constraints)
	if len(failedConstraints) > 0 {
		fmt.Printf("Validation failed, the following constraints failed on the following entries:\n")
		for key, val := range failedConstraints {
			fmt.Printf("Constraint name: [%s]:\n", key)
			for _, v := range val {
				fmt.Printf("\tFailed - Credential name: [%s]\n", v)
			}
			fmt.Println("")
		}
	}
}
