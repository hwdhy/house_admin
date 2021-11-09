package user

import (
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HouseSubscribeDeleteInput struct {
	SubscribeId int `json:"subscribeId"`
}

//用户取消预约看房
func HouseSubscribeDelete(c *gin.Context) {
	var input HouseSubscribeDeleteInput
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}

	userId := utilsToken.GetContextUserId(c)

	db.DB.Model(models.HouseSubscribe{}).
		Where("id = ?", input.SubscribeId).Where("user_id = ?", userId).Delete(&models.HouseSubscribe{})

	c.JSON(http.StatusOK, "取消预约成功")
}
