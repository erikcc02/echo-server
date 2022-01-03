package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	GrayColor   = "\033[90m%s\033[0m"
	YellowColor = "\033[1;33m%s\033[0m"
)

func main() {
	echoHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Printf(GrayColor, "starting request")

		extractHeaders(w, req)
		io.WriteString(w, "teste")

		log.Printf(GrayColor, "end request")
	}

	http.HandleFunc("/", echoHandler)
	log.Println("Listing for requests at http:localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func extractHeaders(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "%q: %q\n", k, v)
		fmt.Printf(YellowColor, fmt.Sprintf("%s: %s\n", k, v))
	}
}
