package main

import "fmt"

func main() {
	a := 5
	b := 6

	fmt.Println("a, b: ", a, b)

	a = a + b // 11
	b = a - b // 11 - 6 = 5
	a = a - b // 11 - 5 = 6

	fmt.Println("a, b: ", a, b)
}