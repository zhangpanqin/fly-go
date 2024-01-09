package main

import "fmt"

type Code string

const (
	LuckyDog Code = "LuckyDog1"
	WatchDog Code = "WatchDog1"
)

type Gender int

const (
	boy Gender = iota
	girl
)

func main() {
	strings := make(map[Code]string)
	strings[LuckyDog] = "111"
	strings[LuckyDog] = "112"
	strings[WatchDog] = "113"
	fmt.Println("map:", strings)
	m := make(map[Gender]string)
	m[boy] = "nan"
	m[girl] = "nv"
	fmt.Println("map:", m)
}
