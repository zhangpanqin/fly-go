package main

import "fmt"

func main() {
	var i = "20"
	fmt.Println(i)
	p := &i
	*p = "1"
	fmt.Println(i)
}
