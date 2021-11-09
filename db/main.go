package db

import (
	"gin_test/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "root:dhyy@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(db.Error)
	}

	RegisterTable(db)
	DB = db
}

func RegisterTable(db *gorm.DB) {
	err := db.AutoMigrate(
		models.BsCity{},
		models.BsArea{},
		models.House{},
		models.HouseDetail{},
		models.HousePicture{},
		models.HouseStar{},
		models.HouseSubscribe{},
		models.HouseTag{},
		models.SupportAddress{},
		models.User{},
	)

	if err != nil {
		log.Fatal(err)
	}
}
