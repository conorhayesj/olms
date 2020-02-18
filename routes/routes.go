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
	r.HandleFunc("/api/v1/book/{bookId}", middleware.GetBook).Methods("GET")
	r.HandleFunc("/api/v1/search/{bookSearch}", middleware.GetBookByName).Methods("GET")
	r.HandleFunc("/api/v1/addBook", middleware.AddBook).Methods("PUT")
	r.HandleFunc("/api/v1/delBook/{bookId}", middleware.DeleteBook).Methods("DELETE")
	return r
}

func home(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/allBooks.html"))
        tmpl.Execute(w, middleware.GetAllBooks(w, r))
}
