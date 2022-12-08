package db

import "gorm.io/gorm"

type Img struct {
	gorm.Model
	Base64 string `gorm:"type:blob"`
}
