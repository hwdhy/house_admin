package db

import (
	"gin_test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDatabase() {
	//连接数据库175.27.224.55
	dsn := "host=localhost user=postgres password=123456 dbname=changfang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
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
