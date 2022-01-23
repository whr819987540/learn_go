// count the duplicate line in different files

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// file names from command params
	if len(os.Args) == 1 {
		fmt.Println("no file name input")
		return
	}
	line_num_map := make(map[string]int)
	// open, get the whole content of a file and split into lines
	for _, file_name := range os.Args[1:] {
		data, err := ioutil.ReadFile(file_name) // return byte slice
		if err != nil {
			fmt.Printf("%s open error, %v", file_name, err)
			continue
		}
		// split the whole content into lines
		for _, line := range strings.Split(string(data), "\n") {
			line_num_map[line]++
		}
	}
	// output if num greater than 1
	for line, num := range line_num_map {
		if num > 1 {
			fmt.Println(line, num)
		}
	}
}
