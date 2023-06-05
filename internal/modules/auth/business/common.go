package business

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	AccessSecretKey     = "ACCESS_SECRET_KEY"
	RefreshSecretKey    = "REFRESH_SECRET_KEY"
	AccessTokenExpired  = 1
	RefreshTokenExpired = 8
)

func GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func VerifyPassword(hashPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func GenerateToken(authorized bool, userID int, name, email, secretKey string, expireInHours int) (string, error) {
	claims := jwt.MapClaims{}

	now := time.Now()

	claims["authorized"] = authorized
	claims["iss"] = "https://my-domain.auth0.com"
	claims["sub"] = "auth0|123456"
	claims["name"] = name
	claims["email"] = email
	claims["user_id"] = userID
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(time.Hour * time.Duration(expireInHours)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(token, secretKey string) {
	secret := []byte(secretKey)

	decodedToken, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra loại phương thức ký
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Trả về secret key
		return secret, nil
	})
	// Kiểm tra lỗi
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}
	// Lấy claims từ token
	if claims, ok := decodedToken.Claims.(*jwt.MapClaims); ok && decodedToken.Valid {
		fmt.Println("Hello,", (*claims)["user_id"])
		fmt.Println("Token expires at:", time.Unix(int64((*claims)["exp"].(float64)), 0))
	} else {
		fmt.Println("Invalid token")
	}
}

func GenerateIDToken() (string, error) {
	return "", nil
}
