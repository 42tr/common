package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:123456@tcp(192.168.1.2:3306)/common?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Img{})
}

func SaveImg(url, base64 string) (int, error) {
	img := Img{Url: url, Base64: base64}
	tx := DB.Create(&img)
	return img.ID, tx.Error
}

func GetImg(id uint) (Img, error) {
	var img Img
	tx := DB.First(&img, id)
	return img, tx.Error
}
