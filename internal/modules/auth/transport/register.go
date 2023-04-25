package transport

import (
	"github.com/gin-gonic/gin"
	"iam/common"
	"iam/internal/modules/auth/business"
	"iam/internal/modules/auth/storage"
	mhttp "iam/sdk/httpserver"
	"net/http"
)

type RegisterBody struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterVerificationBody struct {
	PhoneNumber string `json:"phone_number"`
	OTP         string `json:"otp"`
}

func Register(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body RegisterBody
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		st := storage.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewRegisterBusiness(appCtx, st)

		err := biz.Register(c.Request.Context(), body.PhoneNumber, body.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, mhttp.SimpleSuccessResponse("success"))
	}
}

func VerifyRegister(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body RegisterVerificationBody
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		st := storage.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewRegisterBusiness(appCtx, st)

		data, err := biz.VerifyRegister(c.Request.Context(), body.PhoneNumber, body.OTP)
		if err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, mhttp.SimpleSuccessResponse(data))
	}
}
