package open

import (
	"gin_test/db"
	"gin_test/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PhoneInput struct {
	Phone string `json:"phone"`
}

//手机号验证
func Phone(c *gin.Context) {
	var input PhoneInput
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}

	var user models.User
	db.DB.Model(models.User{}).Where("phone_number = ?", input.Phone).First(&user)
	if user.Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"exist": true,
		})
	}
}
