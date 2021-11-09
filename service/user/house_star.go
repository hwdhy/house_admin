package user

import (
	"fmt"
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HouseStarInput struct {
	HouseId uint `json:"houseId"`
}

//用户收藏房源
func HouseStar(c *gin.Context) {
	var input HouseStarInput
	if err := c.Bind(&input); err != nil {
		fmt.Println("bind err:", err)
		return
	}

	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	var houseStar models.HouseStar
	houseStar.HouseId = input.HouseId
	houseStar.UserId = userId
	houseStar.CreateTime = time.Now()
	houseStar.LastUpdateTime = time.Now()

	if row := db.DB.Model(models.HouseStar{}).Create(&houseStar).RowsAffected; row == 0 {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
}
