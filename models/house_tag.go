package models

// HouseTag 房屋标签
type HouseTag struct {
	Id      uint   `json:"id" gorm:"primaryKey"` //标签ID
	HouseId uint   `json:"houseId"`              //房源ID
	Name    string `json:"name"`                 //标签名称
}

func (HouseTag) TableName() string {
	return "house_tag"
}
