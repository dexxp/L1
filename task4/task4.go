package main

import (
	"sync"
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Количество воркеров
	N := 4

	var wg sync.WaitGroup

	ch := make(chan int)

	// Цикл создающих воркеров, которые будут читать значения из ch
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(&wg, i, ch)
	}

	j := 0
	// создаём канал для обработки Ctrl + C
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM)
	loop:
	for {
		select {
		// Если в канал terminate что-то пришло, то закрываем каналы и выходим из цикла
		case <-terminate:
			close(ch)
			close(terminate)
			break loop
		default:
			ch <- j
			j++
			time.Sleep(time.Second)			
		}
	}

	// дожидаемся выполнения всех горутин
	wg.Wait()
}

func worker(wg *sync.WaitGroup, numberWorker int, ch chan int) {
	defer wg.Done()

	// пробегаемся по значениям из канала ch и выводим их на экран
	for msg := range ch {
		fmt.Printf("Received message: %d; Worker number: %d\n", msg, numberWorker)
	}
}