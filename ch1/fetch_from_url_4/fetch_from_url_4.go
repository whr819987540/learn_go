// parallel using goroutine

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	prefix := "https://"
	start := time.Now()
	// make a channel to pass string
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		// start a goroutine to get the url
		go function(url, ch)
	}
	// use neither of the return value
	// cause one goroutine will only send one string
	for range os.Args[1:] {
		// blocking receive from channel
		ret := <-ch
		fmt.Println(ret)
	}
	fmt.Printf("%.2fs", time.Since(start).Seconds())
}

// http get, pass the time to main goroutine
func function(url string, ch chan string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		// pass error msg to main goroutine
		// avoid confusing output when both goroutine try to output to stdout
		// uniform output port
		ch <- fmt.Sprintf("%v", err)
		return
	}
	// get the byte length
	byte_num, err := io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		ch <- fmt.Sprintf("%s reading error, %v", url, err)
		return
	}
	// release resource
	response.Body.Close()
	// time and byte length
	ch <- fmt.Sprintf("%.2fs \t%dbytes", time.Since(start).Seconds(), byte_num)
}
