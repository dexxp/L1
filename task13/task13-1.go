package main

import "fmt"

func main() {
	a := 5 // 101
	b := 6 // 110

	fmt.Println("a, b: ", a, b)
	// меняем значения переменных местами
	a = a ^ b // 101 ^ 110 = 011 = 3
	b = a ^ b // 011 ^ 110 = 101 = 5
	a = a ^ b // 011 ^ 101 = 110 = 6

	fmt.Println("a, b: ", a, b)
}