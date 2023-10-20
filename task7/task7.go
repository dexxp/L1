package main

import (
	"sync"
	"fmt"
)

func main() {
	m := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	// конкурентно записываем в мапу 10 значений
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			mu.Lock() // блокируем доступ
			m[i] = i // записываем данные в мапу
			mu.Unlock() // разблокируем доступ
		}(i)
	}

	wg.Wait()

	// выводим мапу
	fmt.Println(m)
}
