package middleware

import (
	"encoding/json"
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
	w.Header().Set("Context-Type", "application/w-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	db.Create(&book)
	json.NewEncoder(w).Encode(book)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/w-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	bookId := mux.Vars(r)
	key := bookId["bookId"]
	book := Book{}
	db.Where("id = ?", key).Find(&book)
	json.NewEncoder(w).Encode(book)
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
        w.Header().Set("Context-Type", "application/w-www-form-urlencoded")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	bookId := mux.Vars(r)
	key := bookId["bookId"]
	book := Book{}
	db.Where("id = ?", key).Find(&book)
	db.Unscoped().Delete(&book)
	json.NewEncoder(w).Encode(book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/w-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	books := []Book{}
	db.Find(&books)
	fmt.Println(books)
	for _, book := range books {
		fmt.Println(book.Name)
	}
	json.NewEncoder(w).Encode(books)
}

func CheckOutBook() {
}
