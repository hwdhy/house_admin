package open

import (
	"crypto/md5"
	"fmt"
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type RegisterByPhoneInput struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

//通过手机号注册
func RegisterByPhone(c *gin.Context) {
	var inputData RegisterByPhoneInput
	err := c.Bind(&inputData)
	if err != nil {
		log.Println("bind err:", err)
		return
	}

	var user models.User
	user.PhoneNumber = inputData.PhoneNumber
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(inputData.Password)))
	user.State = 1
	user.CreateTime = time.Now()
	user.LastLoginTime = time.Now()
	user.LastUpdateTime = time.Now()

	db.DB.Model(models.User{}).Create(&user)

	c.JSON(http.StatusOK,gin.H{
		"token":  utilsToken.Create(user.Id),
	})
}
