package main

import (
	"fmt"
	"sort"
)

func binSearch(numbers []int, searchValue int) int {
	l := 0
	r := len(numbers) - 1

	for l <= r {
		mid := (l + r) / 2

		if (numbers[mid] == searchValue) {
			return mid
		}

		if (searchValue < numbers[mid]) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
	
}

func main() {
	numbers := []int{1, 3, 5, 4, 7, 2}
	numbers1 := []int{1, 3, 5, 7, 8}

	// сортируем числа по возрастанию
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	fmt.Println(binSearch(numbers, 1))
	fmt.Println(binSearch(numbers, 2))
	fmt.Println(binSearch(numbers, 3))
	fmt.Println(binSearch(numbers, 4))
	fmt.Println(binSearch(numbers, 5))
	fmt.Println(binSearch(numbers, 7))
	fmt.Println(binSearch(numbers, -314))

	fmt.Println(binSearch(numbers1, 1))
	fmt.Println(binSearch(numbers1, 3))
	fmt.Println(binSearch(numbers1, 5))
	fmt.Println(binSearch(numbers1, 7))
	fmt.Println(binSearch(numbers1, 8))
	fmt.Println(binSearch(numbers1, 120))
}