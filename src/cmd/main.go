package main

import (
	"log"
	"net/http"

	"github.com/erikcc02/echo-server/src/internal/echo"
)

func main() {
	http.HandleFunc("/", echo.BuildHandler())
	log.Println("Listing for requests at http:localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
