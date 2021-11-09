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

type HouseUpdateInput struct {
	Id           string        `json:"id"`           //主键ID
	Address      string        `json:"address"`      //地址
	CityEnName   string        `json:"cityEnName"`   //城市英文简称
	RegionEnName string        `json:"regionEnName"` //区县英文简称
	Street       string        `json:"street"`       //详细地址
	Floor        int           `json:"floor"`        //楼层信息  1:一楼  2:二楼以上    3：独栋   4：独门独院
	Direction    int           `json:"direction"`    //厂房结构  1:标准厂房   2:钢结构    3:其他
	Area         int           `json:"area"`         //面积
	RentWay      int           `json:"rentWay"`      //出租方式  1:整租    2：分租
	Price        int           `json:"price"`        //价格
	Title        string        `json:"title"`        //标题
	Cover        string        `json:"cover"`        //房屋照片封面
	Tags         []string      `json:"tags"`         //标签
	Description  string        `json:"description"`  //描述
	Pictures     []PicturePath `json:"pictures"`     //图片集合

}

//修改房源信息
func HouseUpdate(c *gin.Context) {
	var input HouseUpdateInput
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}

	//1 判断当前房源的创建人
	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	houseId := basic.StringToInt(input.Id)

	//2  判断房屋详情是否存在
	var houseDetail models.HouseDetail
	db.DB.Model(models.HouseDetail{}).
		Where("house_id = ?", houseId).First(&houseDetail)

	//3  更新房屋信息
	db.DB.Model(models.House{}).
		Where("id = ?", houseId).Where("admin_id = ?", userId).Updates(
		map[string]interface{}{
			"city_en_name":   input.CityEnName,
			"region_en_name": input.RegionEnName,
			"street":         input.Street,
			"floor":          input.Floor,
			"direction":      input.Direction,
			"area":           input.Area,
			"price":          input.Price,
			"title":          input.Title,
			"cover":          "http://localhost/" + input.Cover,
			"status":         1,
		})

	//4 更新房屋详情信息
	db.DB.Model(models.HouseDetail{}).Where("id = ?", houseDetail.ID).Updates(
		map[string]interface{}{
			"description": input.Description,
			"rent_way":    input.RentWay,
			"address":     input.Address,
		})

	//5 移除所有照片信息
	db.DB.Model(models.HousePicture{}).Where("house_id = ?", houseId).Delete(&models.HousePicture{})
	//6 获取照片信息 添加到数据库
	if len(input.Pictures) > 0 {
		pics := make([]models.HousePicture, len(input.Pictures))
		for k := range input.Pictures {
			pics[k].HouseId = uint(houseId)
			pics[k].CdnPrefix = "http://localhost/"
			pics[k].Path = input.Pictures[k].Path
		}
		db.DB.Model(models.HousePicture{}).Create(&pics)
	}

	//7 移除所有标签
	db.DB.Model(models.HouseTag{}).Where("house_id = ?", houseId).Delete(&models.HouseTag{})

	//获取标签信息   保存标签到数据库
	if len(input.Tags) > 0 {
		tags := make([]models.HouseTag, len(input.Tags))
		for k := range input.Tags {
			tags[k].HouseId = uint(houseId)
			tags[k].Name = input.Tags[k]
		}
		db.DB.Model(models.HouseTag{}).Create(&tags)
	}
}
