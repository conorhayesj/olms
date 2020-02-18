package middleware

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
	"github.com/gorilla/mux"
	"net/http"
)

type Book struct {
	gorm.Model	`json:"model"`
	Name		string `json:"name"`
	Avail		bool `json:"avail"`
	Due		time.Time `json:"due"`
}

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DB")
	}
	db.AutoMigrate(&Book{})
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	book.Name = r.FormValue("name")
	fmt.Println(book)
	db.Create(&book)
	fmt.Fprintf(w, "Adding new book.")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)
	key := bookId["bookId"]
	book := Book{}
	db.Where("id = ?", key).Find(&book)
	fmt.Fprintf(w, book.Name)
}

func GetBookReturn(w http.ResponseWriter, r *http.Request) Book {
	bookId := mux.Vars(r)
        key := bookId["bookId"]
        book := Book{}
        db.Where("id = ?", key).Find(&book)
	fmt.Println(book)
	return book
}


func FindBookByName(w http.ResponseWriter, r *http.Request) {
	bookSearch := mux.Vars(r)
	search := bookSearch["bookSearch"]
	search = "%" + search + "%"
	fmt.Println(search)
	books := []Book{}
	db.Where("name LIKE ?", search).Find(&books)
	for _, book := range books {
		fmt.Println(book)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)
	key := bookId["bookId"]
	book := Book{}
	db.Where("id = ?", key).Find(&book)
	db.Unscoped().Delete(&book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) []Book {
	fmt.Println("TEST")
	books := []Book{}
	db.Find(&books)
	fmt.Println(books)
	for _, book := range books {
		fmt.Println(book.Name)
	}
	return books
}

func CheckOutBook() {
}
