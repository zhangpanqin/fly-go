package main

import "fmt"

func main() {
	fmt.Println(var3Demo333(1))
	fmt.Println(var3Demo333(2))
}

func var3Demo333(i int) (r int) {
	r = 20
	if i == 1 {
		return
	}
	return 200
}
