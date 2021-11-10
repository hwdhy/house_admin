package db

import (
	"gin_test/models"
	"testing"
)

//增加城市
func TestSupportInsert(t *testing.T) {
	InitDatabase()

	//城市
	var city models.SupportAddress
	city.BelongTo = "sz"
	city.EnName = "sz"
	city.CnName = "深圳"
	city.Level = "city"
	city.Code = "440300"
	city.BaiduMapLng = 114.085947
	city.BaiduMapLat = 22.547

	DB.Model(models.SupportAddress{}).Create(&city)
}

//增加区域
func TestRegionsInsert(t *testing.T) {

	InitDatabase()

	regions := []models.SupportAddress{
		{0, "sz", "lhq", "罗湖区", "region", "440303", 114.13116, 22.54836},
		{0, "sz", "ftq", "福田区", "region", "440304", 114.05571, 22.52245},
		{0, "sz", "nsq", "南山区", "region", "440305", 113.93029, 22.53291},
		{0, "sz", "baq", "宝安区", "region", "440306", 113.88311, 22.55371},
		{0, "sz", "lgq", "龙岗区", "region", "440307", 114.24771, 22.71986},
		{0, "sz", "ytq", "盐田区", "region", "440308", 114.23733, 22.5578},
		{0, "sz", "gmxq", "光明新区", "region", "440320", 113.896026, 22.777292},
		{0, "sz", "psxq", "坪山新区", "region", "440321", 114.34637, 22.690529},
		{0, "sz", "dpxq", "大鹏新区", "region", "440322", 114.479901, 22.587862},
		{0, "sz", "lhxq", "龙华新区", "region", "440323", 114.036585, 22.68695},
	}

	DB.Model(models.SupportAddress{}).Create(&regions)
}
