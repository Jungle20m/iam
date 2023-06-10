package api

import (
	"fmt"
	"github.com/spf13/cobra"
	"iam/common"
	"iam/internal/server"
	"iam/sdk/httpserver"
	tracersdk "iam/sdk/tracer"
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

			// Tracer
			tracer := tracersdk.NewTracer(
				tracersdk.WithEnvironment(conf.App.Environment),
				tracersdk.WithAppName(conf.App.AppName),
				tracersdk.WithServiceName(conf.App.ServiceName),
				tracersdk.WithServerName(fmt.Sprintf("%s:%s", conf.Api.HttpHost, conf.Api.HttpPort)),
				tracersdk.WithLanguage(conf.App.Language))
			defer tracer.Flush()
			tracer.AttachJaegerProvider("http://localhost:14268/api/traces")

			httpHandler := server.NewHttpHandler(appCtx)

			server := httpserver.New(httpHandler, httpserver.WithAddress(conf.Api.HttpHost, conf.Api.HttpPort))
			server.Start()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

			<-quit

			server.Shutdown()
		},
	}
}
