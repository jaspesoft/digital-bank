package config

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"os"
)

func LoadEnvironmentVariables() {

	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")))

	env := "dev"
	if os.Getenv("GO_ENV") == "prod" {
		env = "prod"
	}

	secretId := fmt.Sprintf("digital_bank.%s", env)
	fmt.Println("secretId: ", secretId)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretId),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		fmt.Println("Error getting secret value: ", err.Error())
		os.Exit(1)
	}

	var secrets map[string]string
	_ = json.Unmarshal([]byte(*result.SecretString), &secrets)

	for key, value := range secrets {
		_ = os.Setenv(key, value)

	}

}
