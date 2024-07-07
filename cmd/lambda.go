package cmd

import (
	"fmt"

	"github.com/kaumnen/aclir/internal/utils"
	"github.com/spf13/cobra"
)

// lambdaCmd represents the lambda command
var lambdaCmd = &cobra.Command{
	Use:   "lambda",
	Short: "Use this command to list all lambda functions in your AWS account.",
	Long:  `This command will list all lambda functions in every region in your AWS account.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lambda called")
		utils.GetAllFunctions()
	},
}

func init() {
	rootCmd.AddCommand(lambdaCmd)
}
