package models

import "time"

// User 用户
type User struct {
	Id             uint      `json:"id" gorm:"primaryKey"` //用户ID
	Name           string    `json:"name"`                 //用户名
	Email          string    `json:"email"`                //邮箱
	PhoneNumber    string    `json:"phoneNumber"`          //手机号
	Password       string    `json:"password"`             //加密后的密码
	State          uint8     `json:"state"`                //用户状态 （1正常 2禁用）
	CreateTime     time.Time `json:"createTime"`           //创建时间
	LastLoginTime  time.Time `json:"lastLoginTime"`        //最近登录时间
	LastUpdateTime time.Time `json:"lastUpdateTime"`       //最近更新时间
	Avatar         string    `json:"avatar"`               //头像
	NickName       string    `json:"nickName"`             //昵称
	Introduction   string    `json:"introduction"`         //个人介绍
}

func (User) TableName() string {
	return "user"
}
