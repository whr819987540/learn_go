package main

import "fmt"

func main() {
	// for initialization; condition; post{
	// statement
	// }
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	// without initialization part
	// for ; condition; post{
	// statement
	// }
	i := 15
	for ; i < 20; i++ {
		fmt.Println(i)
	}

	// without post part
	// for initialization ; condition ; {
	// statement
	// }
	for j := 30; j < 34; {
		fmt.Println(j + 1)
		j++
	}

	// without initialization and post part
	// for condition {
	// statement
	// }
	j := 20
	for j < 25 {
		fmt.Println(j + 1)
		j++
	}
}
