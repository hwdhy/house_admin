package admin

import (
	"gin_test/db"
	"gin_test/models"
	"gin_test/utils/basic"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HouseSubscribesInput struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Status   string `json:"status"`
}

type HouseSubscribesOutput struct {
	Total int64                       `json:"total"`
	List  []HouseSubscribesOutputData `json:"list"`
}

type HouseSubscribesOutputData struct {
	HouseSubscribe models.HouseSubscribe `json:"houseSubscribe"` //预约看房数据
	House          models.House          `json:"house"`          //房源信息
	User           models.User           `json:"user"`           //用户信息
}

//admin 预约房源列表
func HouseSubscribes(c *gin.Context) {
	var input HouseSubscribesInput
	err := c.Bind(&input)
	if err != nil {
		log.Fatal("bind err:", err)
		return
	}
	status := basic.StringToInt(input.Status)

	offset := (input.Page - 1) * input.PageSize
	var houseSubscribes []models.HouseSubscribe
	var Total int64
	tx := db.DB.Model(models.HouseSubscribe{})

	if status != 0 {
		tx.Where("status = ?", status)
	}

	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	//获取预约数据
	tx.Debug().Where("admin_id = ?", userId).Count(&Total).Offset(offset).Limit(input.PageSize).Find(&houseSubscribes)

	//通过预约数据房源ID  获取房源信息
	var houseIds []uint
	var userIds []uint
	for k := range houseSubscribes {
		houseIds = append(houseIds, houseSubscribes[k].HouseId)
		userIds = append(userIds, houseSubscribes[k].UserId)
	}

	if len(houseIds) > 0 {
		//获取房源信息
		var houseList []models.House
		db.DB.Model(models.House{}).Where("id in (?)", houseIds).Find(&houseList)

		houseMap := make(map[uint]models.House, len(houseList))
		for k := range houseList {
			houseMap[houseList[k].ID] = houseList[k]
		}

		//获取用户信息
		var userList []models.User
		db.DB.Model(models.User{}).Where("id in (?)", userIds).Find(&userList)

		userMap := make(map[uint]models.User, len(userList))
		for k := range userList {
			userMap[userList[k].Id] = userList[k]
		}

		//拼接返回数据
		list := make([]HouseSubscribesOutputData, len(houseSubscribes))
		for k := range houseSubscribes {
			list[k].HouseSubscribe = houseSubscribes[k]
			list[k].House = houseMap[houseSubscribes[k].HouseId]
			list[k].User = userMap[houseSubscribes[k].UserId]
		}

		c.JSON(http.StatusOK,
			HouseSubscribesOutput{
				Total: Total,
				List:  list,
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": Total,
		"list":  nil,
	})
}
