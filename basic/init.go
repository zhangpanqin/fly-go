package main

import "fmt"

var a, b, c string

func main() {
	if len(a) == 0 {
		a = "a"
		fmt.Println("赋值 a")
	}

	if len(b) == 0 {
		b = "a"
		fmt.Println("赋值 b")
	}

	if len(a) == 0 {
		a = "a"
		fmt.Println("赋值 a")
	}

	if len(b) == 0 {
		a = "a"
		fmt.Println("赋值 b")
	}
	fmt.Println(c)
}
