package user

import (
	"fmt"
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type HouseSubscribeInput struct {
	Phone       string    `json:"phone"`
	HouseId     int       `json:"houseId"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}

//用户预约看房
func HouseSubscribe(c *gin.Context) {
	var input HouseSubscribeInput
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}

	fmt.Println("1111111111")
	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, "预约看房请先登录")
		return
	}
	fmt.Println("222222222222")

	//判断当前用户是否已经预约看房
	var userSubscribe models.HouseSubscribe
	db.DB.Model(models.HouseSubscribe{}).Debug().
		Where("house_id = ?", input.HouseId).
		Where("user_id = ?", userId).First(&userSubscribe)

	if userSubscribe.Id != 0 {
		return
	}


	//查找房屋的管理员
	var house models.House
	db.DB.Model(models.House{}).Debug().
		Select("admin_id").
		Where("id = ?", input.HouseId).
		First(&house)

	houseSubscribe := models.HouseSubscribe{
		HouseId:        uint(input.HouseId),
		UserId:         userId,
		Description:    input.Description,
		Status:         1,
		CreateTime:     time.Now(),
		LastUpdateTime: time.Now(),
		OrderTime:      input.Time,
		Telephone:      input.Phone,
		AdminId:        house.AdminId,
	}

	db.DB.Model(models.HouseSubscribe{}).Debug().Create(&houseSubscribe)

	c.JSON(http.StatusOK, "预约成功")
}
