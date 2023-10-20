package main

import "fmt"

func main() {
	a := 5 
	b := 6 

	fmt.Println("a, b: ", a, b)
	// меняем значения переменных местами
	a, b = b, a

	fmt.Println("a, b: ", a, b)
}