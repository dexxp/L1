package main

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, 
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 12, 43}

	in := make(chan int)
	out := make(chan int)

	done := make(chan bool)

	go func() {
		for {
			v, ok := <-in
			
			if !ok {
				close(out)
				done <- true
				return
			}

			out <- v * 2			
		}
	}()

	go func() {
		for {
			select {
				case result, ok := <-out:
					if !ok {
						return
					}
					fmt.Println(result)
			}
		}
	}()

	for _, num := range numbers {
		in <- num
	}

	close(in)
	<-done
	close(done)
}