package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	if true {
		k, b := 3, 4
		fmt.Println(k, b)
	}

	fmt.Println(a, b)
}
