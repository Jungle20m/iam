package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"iam/common"
	mhttp "iam/sdk/httpserver"
	"net/http"
	"strings"
)

type authHeader struct {
	Token string `header:"Authorization"`
}

type Credential struct {
	UserID   string `json:"working_site_id"`
	ExpireAt string `json:"worker_site_id"`
	IssueAt  string
}

func AuthMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusNetworkAuthenticationRequired, mhttp.AuthRequireErrorResponse(err))
			c.Abort()
			return
		}

		token := strings.Split(h.Token, "Bearer ")
		if len(token) < 2 {
			err := fmt.Errorf("must provide Authorization header with format `Bearer Token`")
			c.JSON(http.StatusBadRequest, mhttp.BadRequestErrorResponse(err, "authentication invalid", "AUTHENTICATION_INVALID"))
			c.Abort()
			return
		}
		bearerToken := token[1]

		decodedToken, err := jwt.ParseWithClaims(bearerToken, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Return secret key
			return []byte(common.AccessSecretKey), nil
		})
		// Kiểm tra lỗi
		if err != nil {
			c.JSON(http.StatusBadRequest, mhttp.BadRequestErrorResponse(err, "token invalid", "TOKEN_INVALID"))
			c.Abort()
			return
		}

		// Get claims from token
		//var credential Credential
		if claims, ok := decodedToken.Claims.(*jwt.MapClaims); ok && decodedToken.Valid {
			fmt.Println("Hello,", (*claims)["user_id"])
			fmt.Printf("claims: %+v\n", claims)
		} else {
			fmt.Println("Invalid token")
		}

		c.Next()
	}
}
