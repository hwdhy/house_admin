package models

import (
	"gorm.io/gorm"
	"time"
)

// House 房源基本信息
type House struct {
	CityEnName   string `json:"cityEnName"`   //城市英文简称
	RegionEnName string `json:"regionEnName"` //区县英文简称
	Street       string `json:"street"`       //详细地址
	Floor        int    `json:"floor"`        //楼层信息  1:一楼  2:二楼以上    3：独栋   4：独门独院
	Direction    int    `json:"direction"`    //厂房结构  1:标准厂房   2:钢结构    3:其他
	Area         int    `json:"area"`         //面积
	Price        int    `json:"price"`        //价格
	Title        string `json:"title"`        //标题
	Cover        string `json:"cover"`        //房屋照片封面

	Status     int  `json:"status"`     //房源状态 （0：未审核  1：审核通过  2：已出租   3；已删除）
	AdminId    uint `json:"adminId"`    //所属管理员ID
	WatchTimes int  `json:"WatchTimes"` //被查看次数

	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (House) TableName() string {
	return "house"
}
