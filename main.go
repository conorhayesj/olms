package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

type Book struct {
	gorm.Model	`json:"model"`
	Name		string `json:"name"`
	Avail		bool `json:"avail"`
	Due		time.Time `json:"due"`
}

var db *gorm.DB
var err error

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", getAllBooks)
	r.HandleFunc("/api/v1/book/{bookId}", getBook).Methods("GET")
	r.HandleFunc("/api/v1/search/{bookSearch}", getBookByName).Methods("GET")
	r.HandleFunc("/api/v1/addBook", addBook).Methods("PUT")
	r.HandleFunc("/api/v1/delBook/{bookId}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wooo. Homepage")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)
	db.Create(&book)
	fmt.Fprintf(w, "Adding new book.")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)
	key := bookId["bookId"]
	book := Book{}
	db.Where("id = ?", key).Find(&book)
	fmt.Fprintf(w, book.Name)
}

func getBookByName(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("searchByName.html"))
	bookSearch := mux.Vars(r)
	search := bookSearch["bookSearch"]
	search = "%" + search + "%"
	fmt.Println(search)
	books := []Book{}
	db.Where("name LIKE ?", search).Find(&books)
	for _, book := range books {
		fmt.Println(book)
	}
	tmpl.Execute(w, books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)
	key := bookId["bookId"]
	book := Book{}
	db.Where("id = ?", key).Find(&book)
	db.Unscoped().Delete(&book)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("allBooks.html"))
	books := []Book{}
	db.Find(&books)
	fmt.Println(books)
	for _, book := range books {
		fmt.Println(book.Name)
	}
	tmpl.Execute(w, books)
}

func checkOutBook() {
}

func main() {
	fmt.Println("sqlite test")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DB")
	}
	defer db.Close()
	db.AutoMigrate(&Book{})

	handleRequests()
}
