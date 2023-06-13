package api

import (
	"fmt"
	"github.com/spf13/cobra"
	"iam/common"
	"iam/config"
	"iam/internal/rpc"
	"iam/internal/server"
	"iam/sdk/httpserver"
	"iam/sdk/mgorm"
	tracersdk "iam/sdk/tracer"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func NewServer(conf *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "run iam api",
		Long:  "run iam api",
		Run: func(cmd *cobra.Command, args []string) {
			// Do something
			fmt.Println("service is started")

			// Database
			db, err := mgorm.New(conf.Mysql.Dsn)
			if err != nil {
				log.Fatalf("connect mgorm error: %v\n", err)
			}

			appCtx := common.NewAppContext(conf, db.Connection)

			// Tracer
			tracer := tracersdk.NewTracer(
				tracersdk.WithEnvironment(conf.App.Environment),
				tracersdk.WithAppName(conf.App.AppName),
				tracersdk.WithServiceName(conf.App.ServiceName),
				tracersdk.WithServerName(fmt.Sprintf("%s:%s", conf.Api.HttpHost, conf.Api.HttpPort)),
				tracersdk.WithLanguage(conf.App.Language))
			defer tracer.Flush()
			tracer.AttachJaegerProvider("http://localhost:14268/api/traces")

			// REST api
			httpHandler := server.NewHttpHandler(appCtx)
			server := httpserver.New(httpHandler, httpserver.WithAddress(conf.Api.HttpHost, conf.Api.HttpPort))
			go func() {
				server.Start()
			}()

			// Grpc
			grpc := rpc.NewServer(appCtx)
			go func() {
				fmt.Println("grpc server is starting...")
				grpc.Serve("", 9090)
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

			<-quit

			server.Shutdown()
		},
	}
}
