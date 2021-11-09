package main

import (
	"gin_test/db"
	"gin_test/service/address"
	"gin_test/service/admin"
	"gin_test/service/house"
	"gin_test/service/open"
	"gin_test/service/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//初始化数据库
	db.InitDatabase()

	r.GET("/open/limits", open.Limits)
	r.POST("/open/login", open.Login)
	r.POST("/open/phone/exist", open.Phone)
	r.POST("/open/registryByPhone", open.RegisterByPhone)

	r.GET("/address/support/cities", address.SupportCities)
	r.GET("/address/support/regions/:enName", address.SupportRegionsList)

	r.POST("/house/houses", house.HouseHouses)
	r.POST("/house/houses/ids", house.HouseByIds)
	r.GET("/house/:houseId", house.HousesById)
	r.POST("/house/search/autocomplete", house.SearchAutocomplete)
	r.POST("/house/map/city/houses", house.MapCityHouses)
	r.POST("/house/map/regions/get", house.MapRegionsGet)

	r.GET("/user", user.User)
	r.POST("/user/logout", user.Logout)
	r.POST("/user/house/star/list", user.HouseStarList)
	r.POST("/user/house/subscribes", user.HouseSubscribes)
	r.POST("/user/house/star", user.HouseStar)
	r.POST("/user/house/unstar", user.HouseUnStar)
	r.POST("/user/upload/photo", user.UploadPhoto)
	r.POST("/user/update", user.Update)
	r.POST("/user/house/subscribe", user.HouseSubscribe)
	r.POST("/user/house/subscribe/delete", user.HouseSubscribeDelete)

	r.POST("/admin/houses", admin.Houses)
	r.POST("/admin/house/upload/photo", admin.UploadPhoto)
	r.POST("/admin/house/add", admin.HouseAdd)
	r.POST("/admin/house/get", admin.HouseGet)
	r.POST("/admin/house/update", admin.HouseUpdate)
	r.POST("/admin/house/subscribes", admin.HouseSubscribes)
	r.POST("/admin/house/operate", admin.HouseOperate)
	r.POST("/admin/house/subscribe/delete", admin.HouseSubscribeDelete)
	r.POST("/admin/house/subscribe/update", admin.HouseSubscribeUpdate)

	r.Run("0.0.0.0:8080")
}
