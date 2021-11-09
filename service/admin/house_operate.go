package admin

import (
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HouseOperateInput struct {
	HouseId uint `json:"houseId"`
	Status  int  `json:"status"`
}

//房源操作
func HouseOperate(c *gin.Context) {
	var input HouseOperateInput
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

	db.DB.Model(models.House{}).
		Where("admin_id = ?", userId).
		Where("id = ?", input.HouseId).
		Update("status", input.Status)
}
