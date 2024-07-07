package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func createEC2Client(ctx context.Context, region string) (*ec2.Client, error) {
	var cfg aws.Config
	var err error

	if region == "" {
		cfg, err = config.LoadDefaultConfig(ctx)
	} else {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithRegion(region))
	}

	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %v", err)
	}

	client := ec2.NewFromConfig(cfg)

	return client, nil
}

func createLambdaClient(ctx context.Context, region string) (*lambda.Client, error) {
	var cfg aws.Config
	var err error

	if region == "" {
		cfg, err = config.LoadDefaultConfig(ctx)
	} else {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithRegion(region))
	}

	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %v", err)
	}

	client := lambda.NewFromConfig(cfg)

	return client, nil
}
