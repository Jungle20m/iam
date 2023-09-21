package http_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"iam/common"
	"iam/internal/modules/iam/business"
	"iam/internal/modules/iam/repository"
	httpsdk "iam/pkg/httpserver"
	tracersdk "iam/pkg/tracer"
)

type registerBody struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func Register(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := tracersdk.NewSpan(c.Request.Context())
		defer span.End()

		var body registerBody
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, httpsdk.HttpErrorResponse(err))
			return
		}

		repo := repository.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewRegisterBusiness(appCtx, repo)

		err := biz.Register(ctx, body.PhoneNumber, body.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, httpsdk.HttpErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, httpsdk.SimpleSuccessResponse("success"))
	}
}

type registerVerificationBody struct {
	ClientID    string `json:"client_id"`
	PhoneNumber string `json:"phone_number"`
	OTP         string `json:"otp"`
}

func VerifyRegister(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body registerVerificationBody
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, httpsdk.HttpErrorResponse(err))
			return
		}

		st := repository.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewRegisterBusiness(appCtx, st)

		data, err := biz.VerifyRegister(c.Request.Context(), body.ClientID, body.PhoneNumber, body.OTP)
		if err != nil {
			c.JSON(http.StatusBadRequest, httpsdk.HttpErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, httpsdk.SimpleSuccessResponse(data))
	}
}
