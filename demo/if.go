package main

import "fmt"

func main() {
	array := []string{"resource", "data", "cpu"}
	array = nil
	if len(array) == 0 {
		fmt.Println("aaa")
	}
}
