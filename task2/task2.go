package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	numbers := [5]float64{2,4,6,8,10}
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func (x float64) {
			defer wg.Done()
			fmt.Println(math.Pow(x, 2))
		}(num)
	}
	
	wg.Wait()
}

