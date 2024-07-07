package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func LambdaFunctionsPerRegion() []string {
	var allFunctions []string

	ctx := context.Background()
	input := RegionInput{
		AllRegions: false,
		DryRun:     false,
	}

	// get enabled regions
	regions, err := GetEnabledRegions(ctx, input)
	if err != nil {
		fmt.Println("Error getting enabled regions:", err)
	} else {
		fmt.Println("Enabled regions:", regions)
	}

	if err != nil {
		fmt.Println(fmt.Errorf("unable to list functions: %v", err))
	}

	for _, region := range regions {
		fmt.Printf("\nLambda functions in region %s:\n", region)

		lambdaClient, err := createLambdaClient(ctx, region)
		if err != nil {
			panic(fmt.Errorf("unable to create EC2 client: %v", err))
		}

		functionsList, err := lambdaClient.ListFunctions(ctx, &lambda.ListFunctionsInput{})

		if err != nil {
			fmt.Println(fmt.Errorf("unable to list functions: %v", err))
		}

		for _, function := range functionsList.Functions {
			fmt.Printf("  - %s\n", aws.ToString(function.FunctionName))
			allFunctions = append(allFunctions, *function.FunctionName)
		}
	}

	return allFunctions
}
