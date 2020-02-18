package main

import (
	"fmt"
	"net/http"
	"./routes"
)

func main() {

	r := routes.Router()

	fmt.Println("sqlite test")

	http.ListenAndServe(":8090", r)
}
