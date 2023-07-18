package main

import "fmt"

const name = 1
const namea1 int = 1

var name1 int = demo222()

func main() {
	fmt.Println(name, name1, namea1)
}

func demo222() int {
	return 200
}

func demo333(i int) (r int) {
	r = 20
	if i == 1 {
		return
	}
	return 200
}
