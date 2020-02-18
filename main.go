package main

import (
	"fmt"
	"net/http"
	"./routes"
)

func main() {

	r := routes.Router()

	fmt.Println("sqlite test")

	http.ListenAndServe(":8080", r)
}
