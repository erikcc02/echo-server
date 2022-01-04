package echo

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	GrayColor   = "\033[90m%s\033[0m"
	YellowColor = "\033[1;33m%s\033[0m"
	BlueColor   = "\033[1;36m%s\033[0m"
)

func BuildHandler() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf(GrayColor, "starting request")

		extractQueryRequest(w, req)
		extractHeaders(w, req)
		extractBody(w, req)

		log.Printf(GrayColor, "end request")
	}
}

func extractQueryRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s --> %s\n", req.Method, req.URL)
	fmt.Printf(BlueColor, fmt.Sprintf("%s --> %s\n", req.Method, req.URL))
}

func extractHeaders(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "%s: %s\n", k, v)
		fmt.Printf(YellowColor, fmt.Sprintf("%s: %s\n", k, v))
	}
}

func extractBody(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	bodyStr := fmt.Sprintf("body:\n %s\n", body)

	fmt.Println(bodyStr)
	io.WriteString(w, bodyStr)
}
