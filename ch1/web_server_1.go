package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle_request)
	http.ListenAndServe("localhost:8000", nil)
}

func handle_request(response http.ResponseWriter, request *http.Request) {
	// write sth to response
	fmt.Fprintf(response, "URL.PATH=%q\n", request.URL.Path)
}
