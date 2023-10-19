package main

import "fmt"

func setBit(n, i int64) int64 {
	return n | int64(1 << i)
}

func clearBit(n, i int64) int64 {
	return n &^ int64(1 << i)
}

func main() {
	n := int64(10) // 0b1010
	index := int64(2) // 0b10
	n = setBit(n, index) // 0b1010 | 0b0100 = 0b1110 = 14
	fmt.Println(n)
	n = clearBit(n, index) // 0b1110 &^ 0b0100 = 0b1010
	fmt.Println(n)
}