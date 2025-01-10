package models

import (
	"github.com/Rakesh-honawad/Book-store/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Book model structure
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}) // Automatically migrate the schema
}

// CreateBook creates a new book record in the database
func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

// GetAllBooks retrieves all book records from the database
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// GetBookById retrieves a book record by its ID
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID = ?", Id).First(&book)
	return &book, db
}

// DeleteBook deletes a book record by its ID
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(&book)
	return book
}
