package main

import (
	"fmt"
)

func main() {
	// два множества чисел
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}


	// мапа для хранения пересечений
	intersection := make(map[int]bool)

	// Добавляем элементы первого множества в мапу
	for _, num := range set1 {
		intersection[num] = true
	}

	// Проверяем элементы второго множества на наличие в карте
	for _, num := range set2 {
		if intersection[num] {
			fmt.Println(num)
		}
	}
}