package transport

import (
	"github.com/gin-gonic/gin"
	"iam/common"
	"iam/internal/modules/auth/business"
	"iam/internal/modules/auth/storage"
	mhttp "iam/sdk/httpserver"
	"net/http"
)

type LoginBody struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func Login(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body LoginBody
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		st := storage.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewLoginBusiness(appCtx, st)

		if err := biz.Login(c.Request.Context(), body.PhoneNumber, body.Password); err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, mhttp.SimpleSuccessResponse(body))
	}
}

func Logout(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, mhttp.SimpleSuccessResponse("success"))
	}
}
