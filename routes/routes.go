package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

func handleRequests() {
        r := mux.NewRouter()
	r.HandleFunc("/", getAllBooks)
	r.HandleFunc("/api/v1/book/{bookId}", getBook).Methods("GET")
	r.HandleFunc("/api/v1/search/{bookSearch}", getBookByName).Methods("GET")
	r.HandleFunc("/api/v1/addBook", addBook).Methods("PUT")
	r.HandleFunc("/api/v1/delBook/{bookId}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
