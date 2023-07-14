package main

import "fmt"

type DemoObject struct {
	Name string
}

func main() {
	obj := DemoObject{
		Name: "xiaoming",
	}
	demo3(&obj)
	fmt.Println(obj)
}

func demo2(d DemoObject) {
	d.Name = "demo2"
	fmt.Println(d)
}

func demo3(d *DemoObject) {
	d.Name = "demo3"
	fmt.Println(*d)
}

func demo() {
	i := 42
	var p *int
	p = &i

	fmt.Println(*p) // 42
	fmt.Println(p)  // 0xc00011c018
	fmt.Println(i)  // 42

	*p = 21 // 修改指针就相当于修改 i

	fmt.Println(*p) // 21
	fmt.Println(p)  // 0xc00011c018
	fmt.Println(i)  // 21

	// 修改 i，同样影响指针
	i = 404
	fmt.Println(*p) // 404
	fmt.Println(p)  // 0xc00011c018
	fmt.Println(i)  // 404
}
