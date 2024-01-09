package main

import (
	"fmt"
)

func main() {
	a := 1
	if true {
		a := 3
		fmt.Println(a)
	}

	fmt.Println(a)
}
