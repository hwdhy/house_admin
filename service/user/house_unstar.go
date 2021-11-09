package user

import (
	"fmt"
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HouseUnStarInput struct {
	HouseId uint `json:"houseId"`
}

//用户取消房源收藏
func HouseUnStar(c *gin.Context) {
	var input HouseUnStarInput
	if err := c.Bind(&input); err != nil {
		fmt.Println("bind err:", err)
		return
	}

	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	db.DB.Debug().
		Where("house_id = ?", input.HouseId).Where("user_id = ?", userId).Delete(&models.HouseStar{})

}
