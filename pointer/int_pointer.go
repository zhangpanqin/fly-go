package main

import "fmt"

func main() {
	var i = 20
	fmt.Println(i)
	p := &i
	fmt.Println(p)
	*p = 1
	fmt.Println(i)
}
