package main

import "fmt"

type People interface {
	Eate(name string)
}

type Student struct {
	Age int
}

func (s Student) Eate(name string) {
	fmt.Printf("main s 指针：%d", s)
	s.Age = 12
	fmt.Printf("%s,Eate", name)
}

func main() {
	s := Student{
		Age: 11,
	}
	s_pointer := &s

	fmt.Println(*s_pointer)

	s.Eate("xiaoming")

	fmt.Printf("调用之后，main s 指针：%d", s)
}
