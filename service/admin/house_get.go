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

type HouseGetInput struct {
	HouseId string `json:"houseId"`
}

type HouseGetOutput struct {
	House HouseGetOutputData `json:"house"`
	City  CityEnName         `json:"city"`
}

type HouseGetOutputData struct {
	models.House                           //房源信息
	HousePictureList []models.HousePicture `json:"housePictureList"` //房源图片
	HouseDetail      models.HouseDetail    `json:"houseDetail"`      //房源详细信息
	Tags             []string              `json:"tags"`             //房源标签
}

type CityEnName struct {
	EnName string `json:"enName"`
}

//房源信息获取
func HouseGet(c *gin.Context) {
	var input HouseGetInput
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}
	houseId := basic.StringToInt(input.HouseId)

	userId := utilsToken.GetContextUserId(c)

	var houseData models.House
	db.DB.Debug().Model(models.House{}).Where("id = ?", houseId).Where("admin_id = ?", userId).First(&houseData)

	var pictureList []models.HousePicture
	db.DB.Debug().Model(models.HousePicture{}).Where("house_id = ?", houseId).Find(&pictureList)

	var houseDetail models.HouseDetail
	db.DB.Debug().Model(models.HouseDetail{}).Where("house_id = ?", houseId).First(&houseDetail)

	var tags []string
	db.DB.Debug().Model(models.HouseTag{}).Where("house_id = ?", houseId).Pluck("name", &tags)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": HouseGetOutput{
			House: HouseGetOutputData{
				House:            houseData,
				HousePictureList: pictureList,
				HouseDetail:      houseDetail,
				Tags:             tags,
			},
			City: CityEnName{
				EnName: houseData.CityEnName,
			},
		},
	})

}
