package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"iam/cmd/api"
	"iam/config"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "run IAM service",
	Long:  "run IAM service",
}

func init() {
	// Init dependence
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config error: %v\n", err)
	}

	// Addition command
	apiCmd := api.NewServer(conf)
	rootCmd.AddCommand(apiCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
