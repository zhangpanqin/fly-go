package main

import "fmt"

func main() {
	G(1)
	fmt.Println("在 main 方法执行 test 方法之后，打印")
}

func G(i int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生了 panic：", err)
		}
	}()
	F(i)
}

func G2(i int) {
	defer func() {
		fmt.Println("G2 panic 2")
	}()
	defer func() {
		fmt.Println("G2 panic 1")
	}()
	F(i)
}

func F(i int) {
	defer func() {
		// panic 执行之后，这个代码会被调用
		fmt.Println("panic 执行之后可以执行2")
	}()
	defer func() {
		// panic 执行之后，这个代码会被调用
		fmt.Println("panic 执行之后可以执行1")
	}()
	if i == 1 {
		panic("Something went wrong!")
	}
	fmt.Println("test 没有发生错误") // 当 panic执行之后，这行代码不会执行
}
