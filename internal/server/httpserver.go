package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"iam/common"
	"iam/docs"
	"iam/internal/modules/auth/transport"
	"net/http"
)

func NewHttpHandler(appCtx common.IAppContext) *gin.Engine {
	config := appCtx.GetConfig()

	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	//// Api logger
	//gin.DisableConsoleColor()
	//f, _ := os.Create(filepath.Join(config.Log.Folder, config.Log.ApiLogFile))
	//gin.DefaultWriter = io.MultiWriter(f)
	//handler.Use(gin.Logger())

	handler.Use(gin.Recovery())

	// Swagger
	docs.SwaggerInfo.Title = config.Swagger.Title
	docs.SwaggerInfo.Version = config.Swagger.Version
	docs.SwaggerInfo.Host = config.Swagger.Host
	docs.SwaggerInfo.BasePath = "/iam/v1"
	handler.GET("/iam/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//// Authentication
	//handler.Use(middleware.AuthMW())

	// Router
	v1 := handler.Group("iam/v1/")

	// health check
	v1.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "pong") })

	v1.POST("/register", transport.Register(appCtx))
	v1.POST("/register/verification", transport.VerifyRegister(appCtx))
	v1.POST("/login", transport.Login(appCtx))
	v1.POST("/logout", transport.Logout(appCtx))
	v1.POST("/reset-password")
	v1.POST("/forget-password")

	return handler

}
