package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aws-secrets-dotenv",
	Short: "A AWS Secrets Manager json secret to dotenv tool",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().String("aws-region", "us-east-2", "Specify the aws region")
}

// Execute application commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		showErrorAndExit(err)
	}
}

func showErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func getFlagStringOrFailed(cmd *cobra.Command, name string) string {
	value, err := cmd.Flags().GetString(name)

	if err != nil {
		showErrorAndExit(fmt.Errorf("Get \"%s\" flag falied: %w", name, err))
	}

	return value
}
