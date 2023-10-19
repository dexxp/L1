package main

import (
	"fmt"
	"time"
)

func main() {
	N := time.Duration(5)

	ch := make(chan int)

	go func() {
		for {
			v := <-ch
			fmt.Println("Received: ", v)
		}
	}()

	go func() {
		for i := 1; ; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	select {
	case <-time.After(N * time.Second):
		close(ch)
		fmt.Println("Exit!")
	}

}