package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	str = strings.ToLower(str)

	// Создаем мапу для отслеживания уникальных символов
	charMap := make(map[rune]bool)

	// Итерируемся по каждому символу в строке
	for _, char := range str {
		// Если символ уже есть в мапе, значит он не уникален
		if charMap[char] {
			return false
		}

		// Добавляем символ в мапу
		charMap[char] = true
	}

	return true
}

func main() {
	str1 := "abcd"
	fmt.Printf("%s: %v\n", str1, isUnique(str1))

	str2 := "abCdefAaf"
	fmt.Printf("%s: %v\n", str2, isUnique(str2))

	str3 := "aabcd"
	fmt.Printf("%s: %v\n", str3, isUnique(str3))
}