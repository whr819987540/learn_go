// display the concrete http request info

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/show/", display_info_handler)
	http.ListenAndServe("localhost:8000", nil)
}

func display_info_handler(response http.ResponseWriter, request *http.Request) {
	// first line of http request: method, url, protocol
	fmt.Fprintf(response, "%s %s %s\n", request.Method, request.URL, request.Proto)

	// header
	// http.Header is a type Header map[string][]string
	// so iterate
	for key, value := range request.Header {
		fmt.Fprintf(response, "%q %q\n", key, value)
	}

	// url params
	// a simplified form, when initializing err in 'if' expression
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(os.Stderr, "request parse error, %q\n", err)
		return
	}
	for q, a := range request.Form {
		fmt.Fprintf(response, "q: %q a: %q\n", q, a)
	}
}
