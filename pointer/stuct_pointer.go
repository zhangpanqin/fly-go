package main

import "fmt"

type Demo struct {
	Name string
}

func main() {
	var i = Demo{
		Name: "qian",
	}
	demoa1(&i)
	fmt.Println(i)
	demoa2(i)
	fmt.Println(i)
}

func demoa1(d *Demo) {
	d.Name = "demoa1"
}
func demoa2(d Demo) {
	d.Name = "demoa2"
}
