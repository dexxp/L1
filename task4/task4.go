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
	N := 4

	var wg sync.WaitGroup

	ch := make(chan int)

	go write(ch)

	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(&wg, i, ch)
	}

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM)

	<-terminate
	close(ch)

	wg.Wait()
}

func write(ch chan int) {
	for i := 0; ; i++ {
		ch <- i

		time.Sleep(time.Second)
	}
}

func worker(wg *sync.WaitGroup, numberWorker int, ch chan int) {
	defer wg.Done()

	for msg := range ch {
		fmt.Printf("Received message: %d; Worker number: %d\n", msg, numberWorker)
	}
}