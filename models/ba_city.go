package models

import "time"

// BsCity 市
type BsCity struct {
	CityId       uint      `json:"CityId" gorm:"primaryKey"` //城市ID
	CityCode     string    `json:"cityCode"`                 //城市编码
	CityName     string    `json:"cityName"`                 //城市名称
	ShortName    string    `json:"shortName"`                //名称简写
	ProvinceCode string    `json:"provinceCode"`             //省份编码
	Lng          string    `json:"lng"`                      //经度
	Lat          string    `json:"lat"`                      //纬度
	Sort         uint      `json:"sort"`                     //排序
	GmtCreate    time.Time `json:"gmtCreate"`                //创建时间
	GmtModified  time.Time `json:"gmtModified"`              //修改时间
	Memo         string    `json:"memo"`
	DataState    uint      `json:"dataState"`
	TenantCode   string    `json:"tenantCode"`
}

func (BsCity) TableName() string {
	return "bs_city"
}
