package house

import (
	"gin_test/db"
	"gin_test/models"
	"gin_test/utils/basic"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HousesByIdOutputData struct {
	House   Outputs     `json:"house"`   //房屋信息
	Star    bool        `json:"star"`    //用户收藏状态
	Reserve bool        `json:"reserve"` //是否已预约
	Agent   models.User `json:"agent"`   //房东信息
}

type Outputs struct {
	models.House
	HouseDetail      models.HouseDetail    `json:"houseDetail"`      //详情
	Tags             []string              `json:"tags"`             //标签
	HousePictureList []models.HousePicture `json:"housePictureList"` //图片列表
	SuggestHouses    []models.House        `json:"suggestHouses"`    //推荐房源  todo 暂用随机获取 order by rand()
	StarNumber       int                   `json:"starNumber"`       //收藏次数

}

// HousesById 房源ID查询房源信息  (房源详情)
func HousesById(c *gin.Context) {
	houseId := basic.StringToInt(c.Param("houseId"))

	var house models.House
	db.DB.Model(models.House{}).Where("id = ?", houseId).First(&house)

	var houseDetail models.HouseDetail
	db.DB.Model(models.HouseDetail{}).Where("house_id = ?", houseId).First(&houseDetail)

	var tags []string
	db.DB.Model(models.HouseTag{}).Where("house_id = ?", houseId).Pluck("name", &tags)

	var pictures []models.HousePicture
	db.DB.Model(models.HousePicture{}).Where("house_id = ?", houseId).Find(&pictures)

	var suggestHouses []models.House
	db.DB.Debug().Model(models.House{}).Order("rand()").Limit(3).Find(&suggestHouses)

	userID := utilsToken.GetContextUserId(c)

	var star bool
	//如果用户没有登录  收藏状态为0
	if userID == 0 {
		star = false
	} else {
		//判断当前用户是否已经收藏
		var houseStar models.HouseStar
		db.DB.Debug().Model(models.HouseStar{}).
			Where("house_id = ?", houseId).
			Where("user_id = ?", userID).
			First(&houseStar)

		if houseStar.Id == 0 {
			star = false
		} else {
			star = true
		}
	}

	var reserve bool
	if userID == 0 {
		reserve = false
	} else {
		var houseReserve models.HouseSubscribe
		db.DB.Model(models.HouseSubscribe{}).Where("house_id = ?", houseId).Where("user_id = ?", userID).First(&houseReserve)
		if houseReserve.Id == 0 {
			reserve = false
		} else {
			reserve = true
		}
	}

	//查询房东信息
	var agent models.User
	db.DB.Model(models.User{}).Where("id = ?", house.AdminId).First(&agent)

	//查看当前房源收藏数量
	var Count int64
	db.DB.Model(models.HouseStar{}).Debug().Where("house_id = ?", houseId).Count(&Count)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": HousesByIdOutputData{
			Star:    star,
			Reserve: reserve,
			House: Outputs{
				House:            house,
				HouseDetail:      houseDetail,
				Tags:             tags,
				HousePictureList: pictures,
				SuggestHouses:    suggestHouses,
				StarNumber:       int(Count),
			},
			Agent: agent,
		},
	})
}
