package main

import "fmt"
import "fly-go/sample"

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	sample.Demo()
	fmt.Println("main")
}
