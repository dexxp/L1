package main

// Разработать программу, которая переворачивает подаваемую на ход строку 
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

import "fmt"

func printRune(run []rune) {
	for _, i := range run {
		fmt.Print(string(i))
	}
	fmt.Println()
}

func reverse(str []rune) []rune {
	n := len(str)
	newStr := make([]rune, n)
	for i := n - 1; i >= 0; i-- {
		newStr = append(newStr, str[i])
	}
	return newStr
}

func reverse2(str []rune) []rune {
	n := len(str)
	newStr := make([]rune, n)
	for i := 0; i < n / 2; i++ {
		newStr[i], newStr[n - 1 - i] = str[n - 1 - i], str[i] 
	}
	return newStr
}

func main() {
	str := []rune("главрыба")
	newStr := reverse(str)
	printRune(newStr)
	newStr = reverse2(newStr)
	printRune(newStr)
}