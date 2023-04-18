package transport

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iam/common"
	"iam/internal/modules/auth/business"
	"iam/internal/modules/auth/storage"
	"net/http"
)

type RegisterBody struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func Register(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body RegisterBody
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		fmt.Printf("body: %v\n", body)

		stor := storage.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewUserBusiness(appCtx, stor)

		err := biz.Register(c.Request.Context())

		if err != nil {
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, "success to register")
	}
}
