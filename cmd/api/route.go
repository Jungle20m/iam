package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"iam/common"
	"iam/docs"
	httpHandler "iam/internal/modules/iam/handler/http"
)

func NewHandler(appCtx common.IAppContext) *gin.Engine {
	config := appCtx.GetConfig()

	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	// // Api logger
	// gin.DisableConsoleColor()
	// f, _ := os.Create(filepath.Join(config.Log.Folder, config.Log.ApiLogFile))
	// gin.DefaultWriter = io.MultiWriter(f)
	// handler.Use(gin.Logger())

	handler.Use(gin.Recovery())

	// Swagger
	docs.SwaggerInfo.Title = config.Swagger.Title
	docs.SwaggerInfo.Version = config.Swagger.Version
	docs.SwaggerInfo.Host = config.Swagger.Host
	docs.SwaggerInfo.BasePath = "/iam/v1"
	handler.GET("/iam/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// // Authentication
	// handler.Use(middleware.AuthMW())

	// Router
	v1 := handler.Group("iam/v1/")

	// health check
	v1.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "pong") })

	v1.POST("/register", httpHandler.Register(appCtx))
	v1.POST("/login", httpHandler.Login(appCtx))
	v1.POST("/logout", httpHandler.Logout(appCtx))
	v1.POST("/password/change")

	// v1.POST("/registration/verify", httpHandler.VerifyRegister(appCtx))
	// v1.POST("/password/recover", httpHandler.PasswordRecovery(appCtx))
	// v1.POST("/password/verify", httpHandler.PasswordVerification(appCtx))

	return handler
}
