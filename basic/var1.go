package main

import "fmt"

func main() {
	var a int
	fmt.Println(a)
	var b, c int
	fmt.Println(b, c)

	var d int = 10
	var e = 10
	fmt.Println(d, e)

	a1 := 100
	fmt.Println(a1)
	a1 = 200
	fmt.Println(a1)

	var a2, a3 int
	fmt.Println(a2, a3)

	var a4, a5 = 11, "100"
	fmt.Println(a4, a5)

	var (
		aa         = 10
		aa1 int    = 10
		aa2 string = "10"
		aa3        = "10"
	)
	fmt.Println(aa, aa1, aa2, aa3)
}
