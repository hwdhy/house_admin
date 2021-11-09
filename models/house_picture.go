package models

import "gorm.io/gorm"

// HousePicture 房源图片
type HousePicture struct {
	ID        uint           `json:"id" gorm:"primaryKey"` //图片ID
	HouseId   uint           `json:"houseId"`              //房源ID
	CdnPrefix string         `json:"cdnPrefix"`            //cdn前缀
	Width     int            `json:"width"`                //宽度
	Height    int            `json:"height"`               //长度
	Location  string         `json:"location"`             //本地
	Path      string         `json:"path"`                 //图片路径
	DeletedAt gorm.DeletedAt `json:"deleted_at"`           //删除时间
}

func (HousePicture) TableName() string {
	return "house_picture"
}
