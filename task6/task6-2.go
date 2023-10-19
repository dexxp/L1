package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; ; i++ {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 400)

			if i == 5 {
				fmt.Println("Done!")
				return
			}
		}
	}()

	wg.Wait()

	fmt.Println("Exit!")
}