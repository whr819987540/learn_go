// '/' get the path
//'/count/' get the http request times, self excluded

package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

var count_mutex sync.Mutex // count mutex
var count int              // count the request times
func main() {
	http.HandleFunc("/", root_handler)
	http.HandleFunc("/count/", count_request_times)
	http.ListenAndServe("localhost:8000", nil)
}

// return the request path and count plus 1
func root_handler(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("in root_handler\n")
	count_mutex.Lock()
	count++
	count_mutex.Unlock()
	fmt.Fprintf(response, "%s URL.PATH:%q\n", os.Args[0], request.URL.Path)
}

// return the request times
func count_request_times(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("in count_handler\n")
	count_mutex.Lock()
	fmt.Fprintf(response, "total times:%d", count)
	count_mutex.Unlock()
}
