package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2,4,6,8,10}
	var wg sync.WaitGroup

	squares := make(chan int)

	go func() {
		defer close(squares)

		for _, num := range numbers {
			squares <- num * num
		}
	}()

	sum := 0
	for square := range squares {
		wg.Add(1)
		go func(sq int) {
			defer wg.Done()
			sum += sq
		}(square)
	}
	wg.Wait()
	
	fmt.Println(sum)

	sum = 0
	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sum += num*num
		}(num)
	}

	wg.Wait()

	fmt.Println(sum)
}

