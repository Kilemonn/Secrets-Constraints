package credential_provider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type AwsProvider struct {
	ctx    context.Context
	cfg    aws.Config
	client *secretsmanager.Client
}

func NewAwsProvider(properties map[string]interface{}) (provider AwsProvider, err error) {
	provider.ctx = context.Background()
	provider.cfg, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Printf("failed to load aws configuration, %s\n", err.Error())
		return
	}

	client := secretsmanager.NewFromConfig(provider.cfg)
	provider.client = client

	return
}

func (p AwsProvider) GetCredentialNames() ([]string, error) {
	names := []string{}
	var nextToken *string = nil
	for {
		listSecrets := &secretsmanager.ListSecretsInput{
			NextToken: nextToken,
		}
		resp, err := p.client.ListSecrets(p.ctx, listSecrets)
		if err != nil {
			return []string{}, err
		}

		for _, s := range resp.SecretList {
			names = append(names, *s.Name)
		}

		if resp.NextToken == nil {
			return names, nil
		} else {
			nextToken = resp.NextToken
		}
	}
}

func (p AwsProvider) GetCredentialWithName(key string) (string, error) {
	getSecretInput := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(key),
		// VersionStage: (defaults to AWSCURRENT if unspecified)
	}

	result, err := p.client.GetSecretValue(p.ctx, getSecretInput)
	if err != nil {
		return "", err
	}
	return *result.SecretString, nil
}

func (p AwsProvider) Shutdown() {

}
