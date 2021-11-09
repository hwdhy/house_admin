package models

import (
	"gorm.io/gorm"
	"time"
)

// HouseSubscribe 预约看房数据
type HouseSubscribe struct {
	Id             uint           `json:"id" gorm:"primaryKey"` //ID
	HouseId        uint           `json:"houseId"`              //房屋ID
	UserId         uint           `json:"userId"`               //用户ID
	Description    string         `json:"description"`          //用户描述
	Status         int            `json:"status"`               //预约状态 1：加入待看名单    2：已约看房时间   3：看房完成
	CreateTime     time.Time      `json:"createTime"`           //创建时间
	LastUpdateTime time.Time      `json:"lastUpdateTime"`       //最近更新时间
	OrderTime      time.Time      `json:"orderTime"`            //预约时间
	Telephone      string         `json:"telephone"`            //手机号码
	AdminId        uint           `json:"adminId"`              //房源发布者ID
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`           //删除
}

func (HouseSubscribe) TableName() string {
	return "house_subscribe"
}
