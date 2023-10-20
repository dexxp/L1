// Реализовать все возможные способы остановки выполнения горутины. 

package main

import (
	"fmt"
	"time"
)

func main() {
	// канал-сингал остановки горутины
	done := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			// когда в канал done приходит значени, выполняется кейс
			case <-done:
				fmt.Println("Done!")
				close(done)
				// завершаем выполнение горутины
				return
			default:
				fmt.Println(i)
				i++
				time.Sleep(time.Second)
			}
		}
	}()

	// ждем 5 секунд
	time.Sleep(time.Second * 5)
	// отправляем в done true, сигнализируя горутине, что нужно завершаться
	done <- true

	fmt.Println("Exit!")

}