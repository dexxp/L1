package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2,4,6,8,10}
	var wg sync.WaitGroup

	// Канал для хранения квадратов
	squares := make(chan int)

	// Первый способ

	// Создаём и запускаем горутину, которая пробегается по массиву
	// и отправляет квадраты элементов в канал squares
	go func() {
		defer close(squares)

		for _, num := range numbers {
			squares <- num * num
		}
	}()

	sum := 0
	// Делаем проход по значениями канала squares и конкурентно рассчитываем
	// сумму квадратов массива, пока в канале не закончатся значения
	for square := range squares {
		wg.Add(1)
		go func(sq int) {
			defer wg.Done()
			sum += sq
		}(square)
	}
	wg.Wait()
	
	fmt.Println(sum)

	// Второй способ

	sum = 0
	var mu sync.Mutex
	// Пробегаемся по значениям массива numbers
	for _, num := range numbers {
		wg.Add(1)
		// конкурентно добавляем в sum квадраты элементов numbers
		go func(num int) {
			defer wg.Done()
			// т.к. sum - общий ресурс, мы должны предоставлять монопольный доступ одной горутины к переменной sum
			mu.Lock()
			sum += num*num 
			mu.Unlock()
		}(num)
	}

	wg.Wait()

	fmt.Println(sum)
}

