package aws

import (
	"context"
	"fmt"

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
	ec2Client, err := createEC2Client(ctx, "")

	if err != nil {
		panic(fmt.Errorf("unable to create EC2 client: %v", err))
	}

	describeRegionsInput := &ec2.DescribeRegionsInput{
		AllRegions:  &input.AllRegions,
		DryRun:      &input.DryRun,
		Filters:     input.Filters,
		RegionNames: input.RegionNames,
	}

	result, err := ec2Client.DescribeRegions(ctx, describeRegionsInput)
	if err != nil {
		return nil, fmt.Errorf("unable to describe regions: %v", err)
	}

	var regionNames []string
	for _, region := range result.Regions {
		regionNames = append(regionNames, *region.RegionName)
	}

	return regionNames, nil
}
