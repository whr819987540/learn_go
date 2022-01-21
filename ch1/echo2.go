// for range

package main

import (
	"fmt"
	"os"
)

func main() {
	for index, param := range os.Args[:] {
		fmt.Println(index, param)
	}
}
