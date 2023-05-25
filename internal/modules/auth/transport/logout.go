package transport

import (
	"github.com/gin-gonic/gin"
	"iam/common"
	"iam/internal/modules/auth/business"
	"iam/internal/modules/auth/storage"
	"iam/internal/server/middleware"
	mhttp "iam/sdk/httpserver"
	"net/http"
)

func Logout(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		credential := middleware.GetCredential(c)

		st := storage.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewLogoutBusiness(appCtx, st)

		err := biz.Logout(c.Request.Context(), credential.UserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, mhttp.SimpleSuccessResponse("success"))
	}
}
