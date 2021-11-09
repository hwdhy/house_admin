package user

import (
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

// User 获取用户信息
func User(c *gin.Context) {
	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	var user models.User
	db.DB.Model(models.User{}).Where("id = ?", userId).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
	})
}
