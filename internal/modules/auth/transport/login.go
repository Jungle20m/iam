package transport

import (
	"github.com/gin-gonic/gin"
	"iam/common"
	"iam/internal/modules/auth/storage"
	mhttp "iam/sdk/httpserver"
	"net/http"
)

import (
	"iam/internal/modules/auth/business"
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

		st := storage.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewLoginBusiness(appCtx, st)

		auth, err := biz.Login(c.Request.Context(), body.ClientID, body.PhoneNumber, body.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, mhttp.SimpleSuccessResponse(auth))
	}
}
