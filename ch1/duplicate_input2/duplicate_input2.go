package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file names passed by the command params
	if len(os.Args) == 1 { // no file names
		fmt.Println("expected file names")
		return
	}

	// open file, read and count
	line_num_map := make(map[string]int)
	for _, file_name := range os.Args[1:] {
		// open the file
		// handler is *os.File
		// err can be captured by os.Stderr
		handler, err := os.Open(file_name)
		if err != nil { // open error
			fmt.Printf("%s open error, %T\n", file_name, err)
		} else {
			CountLines(handler, line_num_map)
			err := handler.Close()
			if err != nil {
				fmt.Printf("%s close error, %T\n", file_name, err)
			}
		}
	}

	// output if num greater than 1
	for line, num := range line_num_map {
		if num > 1 {
			fmt.Printf("'%s' %d\n", line, num)
		}
	}
}

// read and count
func CountLines(handler *os.File, line_num_map map[string]int) {
	input := bufio.NewScanner(handler)
	for input.Scan() {
		line_num_map[input.Text()]++
	}
}
