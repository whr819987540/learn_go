package main

import (
	"fmt"
	"os"
)

func main() {
	var command string
	gap := " "
	for i := 0; i < len(os.Args); i++ {
		command += os.Args[i] + gap
	}
	fmt.Println(command)
}
