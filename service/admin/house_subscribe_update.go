package admin

import (
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HouseSubscribeUpdateInput struct {
	ReserveId uint `json:"reserveId"`
	Status    int  `json:"status"`
}

//预约状态更新
func HouseSubscribeUpdate(c *gin.Context) {
	var input HouseSubscribeUpdateInput
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}

	userID := utilsToken.GetContextUserId(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	db.DB.Model(models.HouseSubscribe{}).
		Where("admin_id = ?", userID).
		Where("id = ?", input.ReserveId).
		Update("status", input.Status)

}
