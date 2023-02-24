package model

import "fmt"

type BookQuery struct {
	ID    int
	ISBN  int64
	Name  string
	Price int
}

// Better 1: Struct + switch
func FindBookByInfo2(info BookQuery) ([]Book, error) {
	var books []Book
	db := GetDB()
	switch {
	case info.ID != 0:
		db = db.Where("id = ?", info.ID)
	case info.ISBN != 0:
		db = db.Where("isbn = ?", info.ISBN)
	case info.Name != "":
		db = db.Where("name = ?", info.Name)
	case info.Price != 0:
		db = db.Where("price = ?", info.Price)
	}
	result := db.Find(&books)
	return books, result.Error
}

//client code
func ClientExample1() {

	books, _ := FindBookByInfo2(BookQuery{
		ISBN: 22,
	})
	fmt.Printf("%+v", books)
}
