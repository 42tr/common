package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("common.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Img{})
}

func SaveImg(base64 string) (uint, error) {
	img := Img{Base64: base64}
	tx := DB.Create(&img)
	return img.ID, tx.Error
}

func GetImg(id uint) (string, error) {
	var img Img
	tx := DB.First(&img, id)
	return img.Base64, tx.Error
}
