package main

import "fmt"

// Структура Human
type Human struct {
	name string
	age int
	weight int
	height int
}

// Метод отображения данных о Human
func (h *Human) DisplayInfo() {
	fmt.Printf("Name: %s\nAge:%d\nWeight: %d\nHeight: %d\n", h.name, h.age, h.weight, h.height)
}

// Структура Action, которая содержит структуру Human и встраивает в себя метод DisplayInfo
type Action struct {
	*Human
}

func main() {
	h := Human{name: "Jack", age:21, weight:60, height:160}
	act := Action{Human: &h}
	h.DisplayInfo()
	act.DisplayInfo()
}