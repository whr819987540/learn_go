package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			html, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			} else {
				fmt.Printf("%s\n", html)
				// release resource
				response.Body.Close()
			}
		}
	}
}
