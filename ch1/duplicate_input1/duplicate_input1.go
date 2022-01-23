package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create a map, with string key and int value
	// map provide o(n) to access and modify
	line_num_map := make(map[string]int)

	// bufio.Scan can read from certain source
	// create a scanner
	input := bufio.NewScanner(os.Stdin)

	// read from os.Stdin
	// remove the /n from os.Stdin
	// return True if not empty
	for input.Scan() {
		// Scan.text() get the input
		if input.Text() == "q" {
			break
		}
		// if the key not exists, the value is initialized to zero
		line_num_map[input.Text()] += 1
	}

	// output if value greater than 1
	// iterate, return the key and value
	for line, num := range line_num_map {
		// format print
		if num > 1 {
			fmt.Printf("line:%s, num:%d\n", line, num)
		}
	}
}
