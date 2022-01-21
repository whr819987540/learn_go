// output the file name if duplicate lines contained
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file names from command params
	if len(os.Args) == 1 {
		fmt.Println("no file name input")
		return
	}
	// iterate the params
	for _, file_name := range os.Args[1:] {
		// open file
		handler, err := os.Open(file_name)
		if err != nil {
			fmt.Printf("%s open error: %v", file_name, err)
			continue
		}
		line_num_map := make(map[string]int)
		input := bufio.NewScanner(handler)
		for input.Scan() {
			line := input.Text()
			line_num_map[line]++
			if line_num_map[line] > 1 && line != "" {
				fmt.Printf("%s contains duplicate line: %s\n", file_name, line)
				break
			}
		}
	}
}
