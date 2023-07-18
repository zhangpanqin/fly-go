package main

import "fmt"

func main() {
	arr := []int{1, 2}
	fmt.Println(var4Demo333(arr))
	fmt.Println(arr)
}

func var4Demo333(i []int) []int {
	return append(i, 3)
}
