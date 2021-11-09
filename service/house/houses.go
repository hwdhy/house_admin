package house

import (
	"gin_test/db"
	"gin_test/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type HousesInputData struct {
	Page          int      `json:"page"`          //当前页面
	PageSize      int      `json:"pageSize"`      //每页数量
	Keyword       string   `json:"keyword"`       //关键词
	CityEnName    string   `json:"cityEnName"`    //城市英文名称
	RegionEnName  string   `json:"regionEnName"`  //区域英文名称
	RentWay       int      `json:"rentWay"`       //类型  1：整租  2：分租  ***暂时没用上
	PriceType     string   `json:"priceType"`     //租金类型
	PriceMin      int      `json:"priceMin"`      //min价格
	PriceMax      int      `json:"priceMax"`      //max价格
	Direction     int      `json:"direction"`     //厂房结构
	OrderBy       string   `json:"orderBy"`       //排序名称
	SortDirection string   `json:"sortDirection"` //升序 降序
	Tags          []string `json:"tags"`          //标签
}

// HouseHouses 房源列表
func HouseHouses(c *gin.Context) {
	var input HousesInputData
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}

	offset := (input.Page - 1) * input.PageSize
	var houses []models.House
	var Total int64
	tx := db.DB.Model(models.House{})
	//关键词查询
	if input.Keyword != "" {
		tx.Where("title like ?", "%"+input.Keyword+"%")
	}
	//城市英文名称
	if input.CityEnName != "" {
		tx.Where("city_en_name = ?", input.CityEnName)
	}
	//区域英文名称
	if input.RegionEnName != "" {
		tx.Where("region_en_name = ?", input.RegionEnName)
	}
	switch input.PriceType {
	case "1":
		tx.Where("price <= 10")
	case "2":
		tx.Where("price >=10 and price <=15")
	case "3":
		tx.Where("price >=15 and price <=20")
	case "4":
		tx.Where("price >=20 and price <=25")
	case "5":
		tx.Where("price >=25 and price <=30")
	case "6":
		tx.Where("price >= 30")
	}

	if input.PriceMin != 0 && input.PriceMax != 0 {
		tx.Where("price >= ? and price <= ?", input.PriceMin, input.PriceMax)
	}

	if input.Direction != 0 {
		tx.Where("direction = ?", input.Direction)
	}

	if input.SortDirection != "" && input.OrderBy != "" {
		sort := strings.ToLower(input.SortDirection)
		switch input.OrderBy {
		case "lastUpdateTime":
			tx.Order("last_update_time " + sort)
		case "price":
			tx.Order("price " + sort)
		case "area":
			tx.Order("area " + sort)
		}
	}

	if len(input.Tags) > 0 {
		//查询有该标签的数据
		var houseIds []uint
		db.DB.Model(models.HouseTag{}).Where("name in (?)", input.Tags).Pluck("house_id", &houseIds)

		if len(houseIds) > 0 {
			tx.Where("id in (?)", houseIds)
		}
	}

	//查询状态为 审核通过类型的房源
	tx.Debug().Where("status = 1").Count(&Total).Offset(offset).Limit(input.PageSize).Find(&houses)

	c.JSON(http.StatusOK, gin.H{
		"total": Total,
		"list":  houses,
	})
}
