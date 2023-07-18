package main

import "fmt"

func main() {
	var i = []string{"1", "2"}
	fmt.Println(i)
	demo1(i)
	fmt.Println(i)
	demo2(&i)
	fmt.Println(i)
}

// 一样可以修改 arr 内容
func demo1(arr []string) {
	arr[0] = "demo1"
}

func demo2(arr *[]string) {
	a := *arr
	a[1] = "demo2"
}
