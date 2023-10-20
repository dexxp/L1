package main

import (
	"fmt"
	"context"
	"time"
)

func main() {
	// создаём контекст с функцией его отмены
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		i := 1
		for {
			select {
			// если была вызвана функция cancel, кейс сработает
			case <-ctx.Done():
				fmt.Println("Ctx done!")
				return
			default:
				fmt.Println(i)
				i++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	// ждем три секунды, а после...
	time.Sleep(time.Second * 3)
	// ...закрываем контекст
	cancel()

	fmt.Println("Exit!")
}