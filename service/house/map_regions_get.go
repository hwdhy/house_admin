package house

import (
	"gin_test/db"
	"gin_test/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MapRegionsGetInput struct {
	CityEnName string `json:"cityEnName"`
}

type MapRegionsGetOutput struct {
	Regions []models.SupportAddress `json:"regions"`
	AggData []HouseBucket           `json:"aggData"`
}

type HouseBucket struct {
	Region string `json:"region"`
	Count  int    `json:"count"`
}

//城市房源聚合
func MapRegionsGet(c *gin.Context) {
	var input MapRegionsGetInput
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}
	//判断城市是否存在
	var supportAddress models.SupportAddress
	db.DB.Model(models.SupportAddress{}).Where("en_name = ?", input.CityEnName).First(&supportAddress)
	if supportAddress.ID == 0 {
		return
	}

	//获取所有区县
	var regions []models.SupportAddress
	db.DB.Model(models.SupportAddress{}).
		Where("belong_to = ?", input.CityEnName).
		Where("level = ?", "region").Find(&regions)

	regionsNameList := make([]string, len(regions))
	for k := range regions {
		regionsNameList[k] = regions[k].EnName
	}

	houseBucket := make([]HouseBucket, len(regions))

	db.DB.Debug().Model(models.House{}).
		Select("region_en_name as region", "count(region_en_name) as count").
		Where("region_en_name in (?)", regionsNameList).
		Group("region_en_name").Scan(&houseBucket)

	c.JSON(http.StatusOK, MapRegionsGetOutput{
		Regions: regions,
		AggData: houseBucket,
	})
}
