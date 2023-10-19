package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 2

	// Проверяем, что индекс находится в допустимом диапазоне
	if index >= 0 && index < len(slice) {
		// Используем синтаксис среза для удаления элемента
		slice = append(slice[:index], slice[index+1:]...)
		fmt.Println("Срез после удаления элемента:", slice)
	} else {
		fmt.Println("Индекс находится за пределами диапазона")
	}

	index = 3

	// Проверяем, что индекс находится в допустимом диапазоне
	if index >= 0 && index < len(slice) {
		// Создаем новый срез с длиной на один элемент меньше
		newSlice := make([]int, len(slice)-1)

		// Копируем элементы до индекса
		copy(newSlice, slice[:index])

		// Копируем элементы после индекса
		copy(newSlice[index:], slice[index+1:])

		fmt.Println("Срез после удаления элемента:", newSlice)
	} else {
		fmt.Println("Индекс находится за пределами диапазона")
	}
}