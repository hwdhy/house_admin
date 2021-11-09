package admin

import (
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HouseSubscribeDeleteInput struct {
	ReserveId uint `json:"reserveId"`
}

//预约看房信息 删除
func HouseSubscribeDelete(c *gin.Context) {
	var input HouseSubscribeDeleteInput
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}

	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	db.DB.Where("admin_id = ?", userId).Where("id = ?", input.ReserveId).Delete(&models.HouseSubscribe{})
}
