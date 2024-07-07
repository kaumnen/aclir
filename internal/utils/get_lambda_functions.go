package utils

import (
	"fmt"

	"github.com/kaumnen/aclir/internal/aws"
)

func GetAllFunctions() {
	regions := aws.LambdaFunctionsPerRegion()

	for _, region := range regions {
		fmt.Println(region)
	}
}
