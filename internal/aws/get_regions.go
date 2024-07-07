package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type RegionInput struct {
	AllRegions  bool
	DryRun      bool
	Filters     []types.Filter
	RegionNames []string
}

func GetEnabledRegions(ctx context.Context, input RegionInput) ([]string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %v", err)
	}

	client := ec2.NewFromConfig(cfg)

	describeRegionsInput := &ec2.DescribeRegionsInput{
		AllRegions:  &input.AllRegions,
		DryRun:      &input.DryRun,
		Filters:     input.Filters,
		RegionNames: input.RegionNames,
	}

	result, err := client.DescribeRegions(ctx, describeRegionsInput)
	if err != nil {
		return nil, fmt.Errorf("unable to describe regions: %v", err)
	}

	var regionNames []string
	for _, region := range result.Regions {
		regionNames = append(regionNames, *region.RegionName)
	}

	return regionNames, nil
}
