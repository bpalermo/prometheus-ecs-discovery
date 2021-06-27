package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// GetParameter
// Retrieve the value for a given SSM Parameter name
func GetParameter(ctx context.Context, parameterName string) *string {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := ssm.NewFromConfig(cfg)

	getParameterOutput, err := svc.GetParameter(ctx, &ssm.GetParameterInput{Name: &parameterName})
	if err != nil {
		log.Println(err)
		return aws.String("")
	}
	return getParameterOutput.Parameter.Value
}
