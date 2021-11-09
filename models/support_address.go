package models

// SupportAddress 支持区域
type SupportAddress struct {
	ID          int64   `json:"id" gorm:"primaryKey"` //地区ID
	BelongTo    string  `json:"belongTo"`             //归属
	EnName      string  `json:"enName"`               //英文缩写
	CnName      string  `json:"cnName"`               //中文名称
	Level       string  `json:"level"`                //地区等级
	Code        string  `json:"code"`                 //地区行政代码
	BaiduMapLng float64 `json:"baiduMapLng"`          //百度地图精度
	BaiduMapLat float64 `json:"baiduMapLat"`          //百度地图纬度
}

func (SupportAddress) TableName() string {
	return "support_address"
}
