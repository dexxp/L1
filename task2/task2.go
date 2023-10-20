package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	// Массив с числами
	numbers := [5]float64{2,4,6,8,10}
	var wg sync.WaitGroup

	for _, num := range numbers {
		// Добавляем к счетчику единицу
		wg.Add(1)
		go func (x float64) {
			// Вычитаем из счётчика единицу
			defer wg.Done()
			fmt.Println(math.Pow(x, 2))
		}(num)
	}
	// ожидаем завершения всех горутин
	wg.Wait()
}

