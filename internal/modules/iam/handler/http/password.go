package http_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"iam/common"
	"iam/internal/modules/iam/business"
	"iam/internal/modules/iam/repository"
	mhttp "iam/pkg/httpserver"
)

type recoveryBody struct {
	PhoneNumber string `json:"phone_number"`
	ClientID    string `json:"client_id"`
}

func PasswordRecovery(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body recoveryBody
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		st := repository.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewPasswordBusiness(appCtx, st)

		err := biz.Recover(c.Request.Context(), body.ClientID, body.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, mhttp.SimpleSuccessResponse("success"))
	}
}

type passwordVerification struct {
	PhoneNumber string `json:"phone_number"`
	ClientID    string `json:"client_id"`
	Password    string `json:"password"`
	Otp         string `json:"otp"`
}

func PasswordVerification(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body passwordVerification
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		st := repository.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewPasswordBusiness(appCtx, st)

		err := biz.Verify(c.Request.Context(), body.ClientID, body.PhoneNumber, body.Password, body.Otp)
		if err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, mhttp.SimpleSuccessResponse("success"))
	}
}
