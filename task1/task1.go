package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
	height int
}

func (h *Human) DisplayInfo() {
	fmt.Printf("Name: %s\nAge:%d\nWeight: %d\nHeight: %d\n", h.name, h.age, h.weight, h.height)
}

type Action struct {
	*Human
}

func main() {
	h := Human{name: "Jack", age:21, weight:60, height:160}
	act := Action{Human: &h}
	h.DisplayInfo()
	act.DisplayInfo()
}