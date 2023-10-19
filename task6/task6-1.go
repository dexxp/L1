package main

import (
	"fmt"
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	stop := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-ctx.Done():
				stop <- true
				fmt.Println("Ctx done!")
				return
			default:
				fmt.Println(i)
				i++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	time.Sleep(time.Second * 3)

	cancel()

	<-stop

	fmt.Println("Exit!")
}