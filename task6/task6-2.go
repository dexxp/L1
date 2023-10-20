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

			// если i == 5 то завершаем горутину
			if i == 5 {
				fmt.Println("Done!")
				return
			}
		}
	}()

	// Дожидаемся выполнения горутины
	wg.Wait()

	fmt.Println("Exit!")
}