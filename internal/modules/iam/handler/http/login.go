package http_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"iam/common"
	"iam/internal/modules/iam/business"
	"iam/internal/modules/iam/repository"
	mhttp "iam/pkg/httpserver"
)

type loginBody struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	ClientID    string `json:"client_id"`
}

func Login(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body loginBody
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		st := repository.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewLoginBusiness(appCtx, st)

		auth, err := biz.Login(c.Request.Context(), body.ClientID, body.PhoneNumber, body.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, mhttp.SimpleSuccessResponse(auth))
	}
}
