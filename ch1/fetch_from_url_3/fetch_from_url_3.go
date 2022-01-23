// check the protocol

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	prefix = "https://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		_, err = io.Copy(os.Stdout, response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			response.Body.Close()
			continue
		}
	}
}
