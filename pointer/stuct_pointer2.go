package main

import "fmt"

type Demo2 struct {
	Name string
}

func main() {
	var i = Demo2{
		Name: "qian",
	}
	a := &i
	fmt.Printf("a 的指针 %#p\n", a)
	i.demo()
	fmt.Println(i)

	i.demo2()
	fmt.Println(i)

	//i.demo2()
}

func (d Demo2) demo() {
	d.Name = "demo"
	fmt.Printf("demo d 的指针 %#p\n", &d)
}

func (d *Demo2) demo2() {
	d.Name = "demo2"
	fmt.Printf("demo2 d 的指针 %#p\n", d)
}
