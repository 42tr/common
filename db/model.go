package db

type Img struct {
	ID     int `gorm:"primarykey"`
	Url    string
	Base64 string `gorm:"type:mediumblob"`
}
