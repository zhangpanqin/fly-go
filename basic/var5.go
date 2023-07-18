package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Eat() {
	fmt.Printf("%s Eat \n", man.Name)
}

func (man *Man) Walk() {
	fmt.Printf("%s Walk \n", man.Name)
}

type SuperMan struct {
	Man
}

func (man *SuperMan) Eat() {
	fmt.Printf("SuperMan %s Eat \n", man.Name)
}

func main() {
	aa1 := SuperMan{}
	aa2 := SuperMan{
		Man{
			Name: "小明",
		},
	}
	aa1.Eat()
	aa2.Eat()
}
