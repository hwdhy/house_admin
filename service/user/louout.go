package user

import (
	"fmt"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
)

type LogoutInput struct {
	Token string `json:"token"`
}

// Logout 用户退出登录
func Logout(c *gin.Context) {
	var input LogoutInput
	err := c.Bind(&input)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}

	utilsToken.Logout(input.Token)
}
