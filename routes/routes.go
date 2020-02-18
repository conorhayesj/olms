package routes

import (
	"github.com/gorilla/mux"
	"../middleware"
	"net/http"
	"html/template"
)

func Router() *mux.Router {
        r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/book/{bookId}", book)
	r.HandleFunc("/api/v1/book/{bookId}", middleware.GetBook).Methods("GET")
	r.HandleFunc("/api/v1/search/{bookSearch}", middleware.FindBookByName).Methods("GET")
	r.HandleFunc("/api/v1/addBook", middleware.AddBook).Methods("POST")
	r.HandleFunc("/api/v1/delBook/{bookId}", middleware.DeleteBook).Methods("DELETE")
	return r
}

func home(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/homePage.html"))
        tmpl.Execute(w, middleware.GetAllBooks(w, r))
}

func book(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/book.html"))
	tmpl.Execute(w, middleware.GetBookReturn(w, r))
}
