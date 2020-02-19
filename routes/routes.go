package routes

import (
	"github.com/gorilla/mux"
	"../middleware"
)

func Router() *mux.Router {
        r := mux.NewRouter()
	r.HandleFunc("/api/v1/book/{bookId}", middleware.GetBook).Methods("GET")
	r.HandleFunc("/api/v1/allBooks", middleware.GetAllBooks).Methods("GET")
	r.HandleFunc("/api/v1/search/{bookSearch}", middleware.FindBookByName).Methods("GET")
	r.HandleFunc("/api/v1/addBook", middleware.AddBook).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/delBook/{bookId}", middleware.DeleteBook).Methods("DELETE")
	return r
}

