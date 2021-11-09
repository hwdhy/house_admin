package utilsToken

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var key = []byte("hwdhy-0426-0125")
var expireTime = 24

// CustomClaims 自定义申明
type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// Create 创建token
func Create(userID uint) string {

	claims := CustomClaims{
		userID,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(30*expireTime)).Unix(), //过期时间20分钟
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return ""
	}
	return tokenStr
}

// Parse 解析token
func Parse(tokenStr string) *jwt.Token {

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil
	}
	if token.Valid {
		return token
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			fmt.Println("token is expired or not valid!")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	return nil
}

// GetContextUserId token转换成用户ID
func GetContextUserId(c *gin.Context) uint {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		log.Println("get cookie err:", err)
		return 0
	}
	return GetUserID(cookie)
}

// GetUserID 获取用户ID
func GetUserID(tokenStr string) uint {
	token := Parse(tokenStr)
	if token == nil {
		return 0
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.UserID
	}
	return 0
}

// Logout 使token失效
func Logout(tokenStr string) {
	token := Parse(tokenStr)
	if token == nil {
		return
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		claims.ExpiresAt = time.Now().Unix()
	}
}
