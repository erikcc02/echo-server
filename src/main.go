package main

import (
	"io"
	"log"
	"net/http"
)

const (
	GrayColor = "\033[90m%s\033[0m"
)

func main() {
	echoHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Printf(GrayColor, "starting request")

		io.WriteString(w, "teste")

		log.Printf(GrayColor, "end request")
	}

	http.HandleFunc("/", echoHandler)
	log.Println("Listing for requests at http:localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
