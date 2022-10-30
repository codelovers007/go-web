package main

import (
	"fmt"
	"net/http"

	"github.com/codelovers007/go-web-1/pkg/handlers"
)

const portNumber = ":8000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("server is starting on port:", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
