package models

import (
	"gorm.io/gorm"
	"time"
)

// HouseStar 用户收藏房源
type HouseStar struct {
	Id             uint           `json:"id" gorm:"primaryKey"` //ID
	HouseId        uint           `json:"houseId"`              //房源ID
	UserId         uint           `json:"userId"`               //用户ID
	CreateTime     time.Time      `json:"createTime"`           //创建时间
	LastUpdateTime time.Time      `json:"lastUpdateTime"`       //最近更新时间
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (HouseStar) TableName() string {
	return "house_star"
}
