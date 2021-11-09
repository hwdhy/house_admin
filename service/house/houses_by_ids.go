package house

import (
	"gin_test/db"
	"gin_test/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HouseByIdsInputData struct {
	HouseIdList []string `json:"houseIdList"`
}

// HouseByIds 房源ID获取房源列表 (最近浏览数据)
func HouseByIds(c *gin.Context) {
	var input HouseByIdsInputData
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}

	var houseData []models.House
	db.DB.Model(models.House{}).Where("id in (?)", input.HouseIdList).Find(&houseData)

	c.JSON(http.StatusOK, gin.H{
		"list": houseData,
	})
}
