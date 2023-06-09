package api

import (
	"fmt"
	"github.com/spf13/cobra"
	"iam/common"
	"iam/internal/server"
	"iam/sdk/httpserver"
	"os"
	"os/signal"
	"syscall"
)

func NewServer(appCtx common.IAppContext) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "run iam api",
		Long:  "run iam api",
		Run: func(cmd *cobra.Command, args []string) {
			// Do something
			fmt.Println("service is started")

			conf := appCtx.GetConfig()

			httpHandler := server.NewHttpHandler(appCtx)

			server := httpserver.New(httpHandler, httpserver.WithAddress(conf.App.HttpHost, conf.App.HttpPort))
			server.Start()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

			<-quit

			server.Shutdown()
		},
	}
}
