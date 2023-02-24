package model

import "fmt"

func GetBookByID(id int) (Book, error) {
	var book Book
	db := GetDB()
	result := db.First(&book, id)
	return book, result.Error
}
func GetBookByISBN(isbn int64) (Book, error) {
	var book Book
	db := GetDB()
	result := db.Where("isbn = ?", isbn).First(&book)
	return book, result.Error
}
func GetBookByCategory(category int) ([]Book, error) {
	var books []Book
	db := GetDB()
	result := db.Where("category = ?", category).Find(&books)
	return books, result.Error
}
func GetBookByPrice(price int) ([]Book, error) {
	var books []Book
	db := GetDB()
	result := db.Where("price = ?", price).Find(&books)
	return books, result.Error
}

//client code
func ClientExample() {
	r, _ := GetBookByPrice(200)
	fmt.Println(r)
}
