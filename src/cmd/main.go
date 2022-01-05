package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/erikcc02/echo-server/src/internal/echo"
)

const ECHOSERVER_PORT string = "ECHOSERVER_PORT"

func main() {

	port := getAppPort(8000)

	http.HandleFunc("/", echo.BuildHandler())
	log.Printf("Listing for requests at http:localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func getAppPort(defaultPort int) int {

	if len(os.Getenv(ECHOSERVER_PORT)) == 0 {
		return defaultPort
	}

	port, err := strconv.Atoi(os.Getenv(ECHOSERVER_PORT))

	if err != nil {
		log.Printf("the application is going up on port %d...\n", defaultPort)
		return defaultPort
	}

	return port
}
