package admin

import (
	"fmt"
	"gin_test/conf"
	"gin_test/db"
	"gin_test/models"
	"gin_test/utils/basic"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HouseAddInput struct {
	Address      string        `json:"address"`      //地址
	CityEnName   string        `json:"cityEnName"`   //城市英文简称
	RegionEnName string        `json:"regionEnName"` //区县英文简称
	Street       string        `json:"street"`       //详细地址
	Floor        int           `json:"floor"`        //楼层信息  1:一楼  2:二楼以上    3：独栋   4：独门独院
	Direction    int           `json:"direction"`    //厂房结构  1:标准厂房   2:钢结构    3:其他
	Area         string        `json:"area"`         //面积
	RentWay      int           `json:"rentWay"`      //出租方式  1:整租    2：分租
	Price        string        `json:"price"`        //价格
	Title        string        `json:"title"`        //标题
	Cover        string        `json:"cover"`        //房屋照片封面
	Tags         []string      `json:"tags"`         //标签
	Description  string        `json:"description"`  //描述
	Pictures     []PicturePath `json:"pictures"`     //图片集合
}

type PicturePath struct {
	Path string `json:"path"`
}

//后台添加房源
func HouseAdd(c *gin.Context) {
	var input HouseAddInput
	err := c.Bind(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

	//1 房源数据校验
	area := basic.StringToInt(input.Area)
	price := basic.StringToInt(input.Price)

	userID := utilsToken.GetContextUserId(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	house := models.House{
		CityEnName:   input.CityEnName,
		RegionEnName: input.RegionEnName,
		Street:       input.Street,
		Floor:        input.Floor,
		Direction:    input.Direction,
		Area:         area,
		Price:        price,
		Title:        input.Title,
		Cover:        conf.CdnPath + input.Cover,
		Status:       1, //默认审核通过
		AdminId:      userID,
	}

	//2 新增房屋信息   house
	db.DB.Model(models.House{}).Create(&house)

	//5 保存房屋详情信息
	db.DB.Model(models.HouseDetail{}).Create(&models.HouseDetail{
		Description: input.Description,
		RentWay:     input.RentWay,
		Address:     input.Address,
		HouseId:     house.ID,
	})

	//6 保存房屋图片消息
	if len(input.Pictures) > 0 {
		pics := make([]models.HousePicture, len(input.Pictures))
		for k := range input.Pictures {
			pics[k].HouseId = house.ID
			pics[k].CdnPrefix = conf.CdnPath
			pics[k].Path = input.Pictures[k].Path
		}

		db.DB.Model(models.HousePicture{}).Create(&pics)
	}

	//7 保存房屋标签消息
	if len(input.Tags) > 0 {
		tags := make([]models.HouseTag, len(input.Tags))
		for k := range input.Tags {
			tags[k].HouseId = house.ID
			tags[k].Name = input.Tags[k]
		}

		db.DB.Model(models.HouseTag{}).Create(&tags)
	}

	c.JSON(http.StatusOK, "发布成功")
}
