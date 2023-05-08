package transport

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iam/common"
	mhttp "iam/sdk/httpserver"
	"net/http"
)

type authHeader struct {
	Token string `header:"Authorization"`
}

func Logout(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("hello anh em")

		c.JSON(http.StatusOK, mhttp.SimpleSuccessResponse("success"))
	}
}
