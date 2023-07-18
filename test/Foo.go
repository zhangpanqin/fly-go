package test

import (
	"os"
)

func init() {
	if a := os.Getenv("a"); len(a) == 0 {
		panic("没有设置 a")
	}
}

func Sum1(i int) int {
	return i + 1
}
