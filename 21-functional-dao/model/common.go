package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	ISBN     int64  `json:"isbn"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category int    `json:"category"`
}

func ConnectDB() *gorm.DB {
	d, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db = d
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	ConnectDB()
	db.AutoMigrate(&Book{})
}

func AddFakeBook() {
	db.Create(&Book{
		ISBN:     22,
		Name:     "Book 22",
		Price:    100,
		Category: 2,
	})

	db.Create(&Book{
		ISBN:     23,
		Name:     "Book 23",
		Price:    200,
		Category: 2,
	})
}
