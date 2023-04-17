package transport

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iam/common"
	"iam/internal/modules/auth/business"
	"iam/internal/modules/auth/storage"
	"net/http"
)

func Register(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		stor := storage.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewUserBusiness(appCtx, stor)

		err := biz.Register(c.Request.Context())

		if err != nil {
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, "success")
	}
}
