package main

import "fmt"

func main() {
	var i = []string{""}
	fmt.Println(i)
	Sum223(i)
	fmt.Println(i)
	Sum22(&i)
	fmt.Println(i)
}

// 一样可以修改 arr 内容
func Sum223(arr []string) {
	arr = append(arr, "Sum22 传的指针")
}

func Sum22(arr *[]string) {
	*arr = append(*arr, "Sum22 传的指针")
}
