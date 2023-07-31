package main

import (
	"errors"
	"fmt"
)

func main() {
	err(nil)
	err(errors.New("111"))
	err(errors.New("111"))

	err2(nil)
	err2(&BusinessError{
		Code:    "111",
		Message: "111",
	})
}

func err(err error) {
	if err != nil {
		fmt.Println("err 发生", err)
	} else {
		fmt.Println("没有 err 发生", err)
	}

}

func err2(err *BusinessError) {
	if err != nil {
		fmt.Println("BusinessError 发生", err)
	} else {
		fmt.Println("没有 BusinessError 发生", err)
	}

}

type BusinessError struct {
	Code    string
	Message string
}

func (err *BusinessError) Error() string {
	return err.Message + "-" + err.Code
}
