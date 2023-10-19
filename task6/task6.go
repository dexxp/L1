// Реализовать все возможные способы остановки выполнения горутины. 

package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-done:
				fmt.Println("Done!")
				close(done)
				return
			default:
				fmt.Println(i)
				i++
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 5)

	done <- true

	fmt.Println("Exit!")

}