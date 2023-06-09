package cmd

import (
	"github.com/spf13/cobra"
	"iam/cmd/api"
	"iam/common"
	"iam/config"
	"iam/sdk/mgorm"
	"log"
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

	db, err := mgorm.New(conf.Mysql.Dsn)
	if err != nil {
		log.Fatalf("connect mgorm error: %v\n", err)
	}

	appCtx := common.NewAppContext(conf, db.Connection)

	// Addition command
	apiCmd := api.NewServer(appCtx)
	rootCmd.AddCommand(apiCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
