package main

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, 
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

import (
	"fmt"
)

func main() {
	// слайс чисел
	numbers := []int{1, 2, 3, 4, 5, 12, 43}

	in := make(chan int)
	out := make(chan int)

	done := make(chan bool)

	go func() {
		for {
			select {
			// если в done пришло значение, то...  
			case <-done:
				// ...закрываем канал out...
				close(out)
				// ...и останавливаем горутину
				return 
			// читаем значение из канала in
			case v := <-in:
				// записываем значение из канала in умноженное на два в канал out
				out <- v * 2
			}		
		}
	}()

	go func() {
		for {
			select {
				// читаем результат из канала out
				case result, ok := <-out:
					// если канал закрыт, то останавливаем горутину
					if !ok {
						return
					}
					// иначе выводим на экран значение out
					fmt.Println(result)
			}
		}
	}()

	// проходимся по слайсу
	for _, num := range numbers {
		// отправляем в канал in значения из слайса
		in <- num
	}

	// отправляем сигнал в done
	done <- true
	close(in)
	close(done)
}