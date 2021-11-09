package user

import (
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// HouseStarList 用户收藏房源列表
func HouseStarList(c *gin.Context) {
	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	//通过用户ID查询收藏房源ID
	var houseStars []models.HouseStar
	db.DB.Debug().Model(models.HouseStar{}).Where("user_id = ?", userId).Find(&houseStars)

	starCreateMap := make(map[uint]time.Time)
	houseIds := make([]uint, len(houseStars))
	for k := range houseStars {
		houseIds[k] = houseStars[k].HouseId
		starCreateMap[houseStars[k].HouseId] = houseStars[k].CreateTime
	}

	//通过房源ID查询房源信息
	var houseList []models.House
	db.DB.Debug().Model(models.House{}).Where("id in (?)", houseIds).Find(&houseList)

	//收藏时间加入返回记录
	for k := range houseList {
		houseList[k].CreatedAt = starCreateMap[houseList[k].ID]
	}

	c.JSON(http.StatusOK, gin.H{
		"list": houseList,
	})
}
