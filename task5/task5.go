package main

import (
	"fmt"
	"time"
)

func main() {
	// Время выполнения программы
	N := time.Duration(5)

	// канал для положительных чисел
	ch := make(chan int)
	// канал для обозначения завершения
	done := make(chan bool)

	// горутина читающая значения из канала ch
	go func() {
		for {
			v, ok := <-ch
			// если канал ch закрыт, то завершаем горутину
			if !ok {
				return
			}
			fmt.Println("Received: ", v)
		}
	}()

	// горутина для записи чисел в канал ch
	go func() {
		for i := 1; ; i++ {
			select {
			// если в канал done пришло значение, то завершаем горутину
			case <-done:
				return
			default:
				ch <- i
			}
		}
	}()

	select {
	// когда пройдёт N секунд...
	case <-time.After(N * time.Second):
		// отправляем значение в done 
		done <- true
		// закрываем канал ch 
		close(ch)
	}

	close(done)
}