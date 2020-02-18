package routes

import (
	"github.com/gorilla/mux"
	"../middleware"
)

func Router() *mux.Router {
        r := mux.NewRouter()
	r.HandleFunc("/", middleware.GetAllBooks)
	r.HandleFunc("/api/v1/book/{bookId}", middleware.GetBook).Methods("GET")
	r.HandleFunc("/api/v1/search/{bookSearch}", middleware.GetBookByName).Methods("GET")
	r.HandleFunc("/api/v1/addBook", middleware.AddBook).Methods("PUT")
	r.HandleFunc("/api/v1/delBook/{bookId}", middleware.DeleteBook).Methods("DELETE")
	return r
}
