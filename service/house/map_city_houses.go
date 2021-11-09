package house

import (
	"gin_test/db"
	"gin_test/models"
	"gin_test/utils/basic"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type MapCityHousesInput struct {
	CityEnName    string   `json:"cityEnName"`    //城市英文缩写
	Page          int      `json:"page"`          //当前页面
	PageSize      int      `json:"pageSize"`      //每页数量
	Direction     string   `json:"direction"`     //朝向
	OrderBy       string   `json:"orderBy"`       //排序
	SortDirection string   `json:"sortDirection"` //升序 降序
	PriceMax      string   `json:"priceMax"`      //价格最大值
	PriceMin      string   `json:"priceMin"`      //价格最小值
	PriceRange    string   `json:"priceRange"`    //价格范围
	RentWay       string   `json:"rentWay"`       //1整租 2分租
	Tags          []string `json:"tags"`          //标签
	Bounds        BoundsLL `json:"bounds"`        //地图经纬度

}

//四个角经纬度
type BoundsLL struct {
	LeftTopLatitude      float64 `json:"leftTopLatitude"`      //左上纬度
	LeftTopLongitude     float64 `json:"leftTopLongitude"`     //左上经度
	RightBottomLatitude  float64 `json:"rightBottomLatitude"`  //右下纬度
	RightBottomLongitude float64 `json:"rightBottomLongitude"` //右下经度

}

type MapCityHouseOutput struct {
	Total int64                    `json:"total"`
	List  []MapCityHouseOutputData `json:"list"`
}

type MapCityHouseOutputData struct {
	models.House
	HouseDetail models.HouseDetail `json:"houseDetail"`
	Tags        []string           `json:"tags"`
}

//地图找房
func MapCityHouses(c *gin.Context) {
	var input MapCityHousesInput
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}
	direction := basic.StringToInt(input.Direction)
	priceMin := basic.StringToInt(input.PriceMin)
	priceMax := basic.StringToInt(input.PriceMax)
	rentWay := basic.StringToInt(input.RentWay)

	offset := (input.Page - 1) * input.PageSize
	var houses []models.House
	var Total int64
	tx := db.DB.Model(models.House{})

	//城市英文名称
	if input.CityEnName != "" {
		tx.Where("city_en_name = ?", input.CityEnName)
	}

	if direction != 0 {
		tx.Where("direction = ?", direction)
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

	if priceMin != 0 {
		tx.Where("price >= ?", priceMin)
	}
	if priceMax != 0 {
		tx.Where("price <= ?", priceMax)
	}
	if rentWay != 0 {
		tx.Where("rent_way = ?", rentWay)
	}
	if len(input.Tags) > 0 {
		//查询有该标签的数据
		var houseIds []uint
		db.DB.Model(models.HouseTag{}).Where("name in (?)", input.Tags).Pluck("house_id", &houseIds)

		if len(houseIds) > 0 {
			tx.Where("id in (?)", houseIds)
		}
	}

	if input.Bounds.LeftTopLatitude != 0 {
		//通过经纬度查询区域内县区   ba_area
		var areaNameList []string
		db.DB.Model(models.BsArea{}).
			Where("lng >= ? and  lng <= ?", input.Bounds.LeftTopLongitude, input.Bounds.RightBottomLongitude).
			Where("lat >= ? and lat <= ?", input.Bounds.RightBottomLatitude, input.Bounds.LeftTopLatitude).
			Pluck("area_name", &areaNameList)

		if len(areaNameList) > 0 {
			//通过区域名称查询 区域内房源ID support_address
			var enName []string
			db.DB.Model(models.SupportAddress{}).Where("cn_name in ?", areaNameList).Pluck("en_name", &enName)

			if len(enName) > 0 {
				var houseIds []uint
				db.DB.Model(models.House{}).Where("region_en_name in ?", enName).Pluck("id", &houseIds)

				tx.Where("id in ?", houseIds)
			} else {
				tx.Where("id = 0")
			}
		} else {
			tx.Where("id = 0")
		}
	}

	//查询状态为 审核通过类型的房源
	tx.Debug().Where("status = 1").Count(&Total).Offset(offset).Limit(input.PageSize).Find(&houses)

	var houseIDs []uint
	for k := range houses {
		houseIDs = append(houseIDs, houses[k].ID)
	}
	//查询房源详情信息
	var houseDetail []models.HouseDetail
	db.DB.Model(models.HouseDetail{}).Where("house_id in (?)", houseIDs).Find(&houseDetail)

	houseDetailMap := make(map[uint]models.HouseDetail)
	for k := range houseDetail {
		houseDetailMap[houseDetail[k].HouseId] = houseDetail[k]
	}

	res := make([]MapCityHouseOutputData, len(houses))
	for k := range houses {
		//查找房源标签
		var tags []string
		db.DB.Model(models.HouseTag{}).Where("house_id = ?", houses[k].ID).Pluck("name", &tags)

		res[k].House = houses[k]
		res[k].HouseDetail = houseDetailMap[houses[k].ID]
		res[k].Tags = tags
	}

	c.JSON(http.StatusOK, MapCityHouseOutput{
		Total: Total,
		List:  res,
	})
}
