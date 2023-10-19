package main

import (
	"sync"
	"fmt"
)

type Counter struct {
	count int
}

func (c *Counter) Inc() {
	c.count++
}

func main() {
	N := 10
	counter := Counter{count: 0}
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()

	fmt.Println("Counter: ", counter.count)

	done := make(chan bool)
	go func() {
		for i := 0; i < N; i++ {
			counter.Inc()
		}
		close(done)
	}()

	<-done

	fmt.Println("Counter: ", counter.count)
}