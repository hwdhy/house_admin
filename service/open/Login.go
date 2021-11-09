package open

import (
	"fmt"
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInput struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Login 登录
func Login(c *gin.Context) {
	var input LoginInput
	err := c.Bind(&input)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}

	var user models.User
	db.DB.Debug().Model(models.User{}).
		Where("phone_number = ?", input.Phone).
		Where("password = md5(?)", input.Password).First(&user)
	if user.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"token":  utilsToken.Create(user.Id),
	})
}
