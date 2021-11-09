package user

import (
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	Agent          models.User           `json:"agent"`          //房东信息
}

// 用户预约看房列表
func HouseSubscribes(c *gin.Context) {
	var input HouseSubscribesInput
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}
	status, err := strconv.Atoi(input.Status)

	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	offset := (input.Page - 1) * input.PageSize
	var subscribes []models.HouseSubscribe
	var Total int64

	//通过用户ID获取预约消息
	db.DB.Model(models.HouseSubscribe{}).Debug().
		Where("user_id = ? and status = ?", userId, status).
		Count(&Total).
		Offset(offset).Limit(input.PageSize).
		Find(&subscribes)

	list := make([]HouseSubscribesOutputData, len(subscribes))
	if Total > 0 {
		//获取所有预约的房源ID和用户ID
		var houseIds []uint
		userIds := []uint{userId}
		for k, _ := range subscribes {
			houseIds = append(houseIds, subscribes[k].HouseId)
			userIds = append(userIds, subscribes[k].AdminId)
		}
		//获取所有预约的房源消息
		var houseList []models.House
		db.DB.Model(models.House{}).Debug().Where("id in (?)", houseIds).Find(&houseList)

		//封装房源消息到map中
		houseMap := make(map[uint]models.House, len(houseList))
		for k, _ := range houseList {
			houseMap[houseList[k].ID] = houseList[k]
		}

		//获取所有用户消息
		var userList []models.User
		db.DB.Model(models.User{}).Debug().Where("id in (?)", userIds).Find(&userList)

		userMap := make(map[uint]models.User, len(userList))
		for k, _ := range userList {
			userMap[userList[k].Id] = userList[k]
		}

		for k, _ := range subscribes {
			list[k].HouseSubscribe = subscribes[k]
			list[k].House = houseMap[subscribes[k].HouseId]
			list[k].User = userMap[subscribes[k].UserId]
			list[k].Agent = userMap[subscribes[k].AdminId]
		}

		c.JSON(http.StatusOK, HouseSubscribesOutput{
			Total: Total,
			List:  list,
		})
	}
}
