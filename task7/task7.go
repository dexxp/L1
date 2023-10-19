package main

import (
	"sync"
	"fmt"
)

func main() {
	m := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(m)
}
