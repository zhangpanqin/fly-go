package main

import "fmt"

type Clients []int

var clients Clients

func New() {
	clients = []int{1, 2, 3}
	fmt.Printf("clients %p\n", clients)
}
func (c *Clients) demo2() {
	d := *c
	d = append(d, 4)
	fmt.Printf("demo2,clients %p,%v\n", d, d)
}
func (c Clients) demo3() {
	d := append(c, 4)
	fmt.Printf("demo3,clients %p,%v\n", d, d)
	fmt.Printf("demo3,clients %p,%v\n", clients, clients)
}

func main() {
	New()
	clients.demo2()
	fmt.Printf("clients %p,%v\n", clients, clients)
	clients.demo3()
	fmt.Printf("clients %p,%v\n", clients, clients)
}
