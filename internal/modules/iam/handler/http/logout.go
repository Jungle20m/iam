package http_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"iam/common"
	mhttp "iam/pkg/httpserver"
)

type logoutBody struct {
	ClientID string `json:"client_id"`
}

func Logout(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// credential := middleware.GetCredential(c)
		//
		// var body logoutBody
		// if err := c.ShouldBind(&body); err != nil {
		// 	c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
		// 	return
		// }
		//
		// st := repository.NewMysqlStorage(appCtx.GetDB())
		// biz := business.NewLogoutBusiness(appCtx, st)
		//
		// err := biz.Logout(c.Request.Context(), credential.UserID, body.ClientID)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, mhttp.HttpErrorResponse(err))
		// 	return
		// }

		c.JSON(http.StatusOK, mhttp.SimpleSuccessResponse("success"))
	}
}
