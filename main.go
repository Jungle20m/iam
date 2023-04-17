package main

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"iam/common"
	"iam/config"
	"iam/internal/server"
	"iam/sdk/httpserver"
	"iam/sdk/mysql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Otp() {
	fmt.Println("iam service running")

	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "IAM",
		AccountName: "0327380000",
	})
	fmt.Printf("secret key: %v\n", key.Secret())
	firstOtp, _ := totp.GenerateCode(key.Secret(), time.Now())
	fmt.Printf("otp lan 1: %v\n", firstOtp)

	time.Sleep(1 * time.Second)

	secondOtp, _ := totp.GenerateCode(key.Secret(), time.Now())
	fmt.Printf("otp lan 2: %v\n", secondOtp)

	time.Sleep(3 * time.Second)

	ok := totp.Validate(firstOtp, key.Secret())

	fmt.Printf("result: %v\n", ok)
}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config error: %v\n", err)
	}

	db, err := mysql.New(conf.Mysql.Dns)
	if err != nil {
		log.Fatalf("connect mysql error: %v\n", err)
	}

	appCtx := common.NewAppContext(conf, db)

	httpHandler := server.NewHttpHandler(appCtx)

	server := httpserver.New(httpHandler, httpserver.WithAddress(conf.App.HttpHost, conf.App.HttpPort))
	server.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	server.Shutdown()
}