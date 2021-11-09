package models

import "time"

// BsArea 区
type BsArea struct {
	AreaId      uint      `json:"areaId" gorm:"primaryKey"` //区ID
	AreaCode    string    `json:"areaCode"`                 //区域编码
	CityCode    string    `json:"cityCode"`                 //城市编码
	AreaName    string    `json:"areaName"`                 //区域名称
	ShortName   string    `json:"shortName"`                //名称简写
	Lng         string    `json:"lng"`                      //经度
	Lat         string    `json:"lat"`                      //纬度
	Sort        uint      `json:"sort"`                     //排序
	GmtCreate   time.Time `json:"gmtCreate"`                //创建时间
	GmtModified time.Time `json:"gmtModified"`              //修改时间
	Memo        string    `json:"memo"`                     //描述
	DataState   int       `json:"dataState"`                //数据类型
	TenantCode  string    `json:"tenantCode"`
}

func (BsArea) TableName() string {
	return "bs_area"
}
