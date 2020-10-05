package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/spf13/cobra"
)

func init() {
	command := &cobra.Command{
		Use:   "get [secret-id]",
		Short: "Retrive AWS Secrets Manager json secret to dotenv file",
		Args:  cobra.ExactArgs(1),
		Run:   getCommnad,
	}

	command.Flags().String("version-id", "", "Specify AWS Secrets Manager Secret Version Id")
	command.Flags().String("version-stage", "", "Specify AWS Secrets Manager Secret Version Stage")
	command.Flags().StringP("file", "f", ".env", "Specify dotenv file path")

	rootCmd.AddCommand(command)
}

func getCommnad(cmd *cobra.Command, args []string) {
	versionID := getFlagStringOrFailed(cmd, "version-id")
	versionStage := getFlagStringOrFailed(cmd, "version-stage")
	fileName := getFlagStringOrFailed(cmd, "file")

	awsSession := secretsmanager.New(session.New(&aws.Config{
		Region: aws.String(getFlagStringOrFailed(cmd, "aws-region")),
	}))
	awsInput := &secretsmanager.GetSecretValueInput{
		SecretId: &args[0],
	}

	if len(versionID) > 0 {
		awsInput.VersionId = &versionID
	}

	if len(versionStage) > 0 {
		awsInput.VersionStage = &versionID
	}

	secretsData, err := awsSession.GetSecretValue(awsInput)

	if err != nil {
		showErrorAndExit(fmt.Errorf("Get secret value falied: %w", err))
	}

	var secrets map[string]string
	if err = json.Unmarshal([]byte(*secretsData.SecretString), &secrets); err != nil {
		showErrorAndExit(fmt.Errorf("Unmarshal secrets json falied: %w", err))
	}

	file, err := os.Create(fileName)
	if err != nil {
		showErrorAndExit(fmt.Errorf("Creating %s falied: %w", fileName, err))
	}

	defer file.Close()
	defer file.Sync()

	for key, value := range secrets {
		if _, err := file.WriteString(fmt.Sprintf("%s=%s\n", key, value)); err != nil {
			showErrorAndExit(fmt.Errorf("Writing dotenv falied: %w", err))
		}
	}
}
