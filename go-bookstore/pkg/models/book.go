package models

import (
	"github.com/Vib-UX/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publications"`
}

func init() { // Initialise the database
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Lets create Methods & functions which will talk to our database
func (b *Book) CreateBook() *Book {
	// NewRecord, Create are functions in gorm so we don't have to write direct queries for mysql, gorm will take care of it
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book { // This will return slice of all books
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
