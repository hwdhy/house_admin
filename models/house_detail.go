package models

// HouseDetail 房屋详细描述
type HouseDetail struct {
	ID           uint   `json:"id" gorm:"primaryKey"` //详情ID
	Description  string `json:"description"`          //房屋描述
	LayoutDesc   string `json:"layoutDesc"`           //户型介绍
	Traffic      string `json:"traffic"`              //交通出行介绍
	RoundService string `json:"roundService"`         //周边配套设施
	RentWay      int    `json:"rentWay"`              //出租方式  1整租  2分租
	Address      string `json:"address"`              //房源地址
	HouseId      uint   `json:"houseId"`              //房屋ID
}

func (HouseDetail) TableName() string {
	return "house_detail"
}
