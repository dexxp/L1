package main 

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func sleep1(seconds int) {
	startTime := time.Now()

	for {
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(startTime)
		if elapsedTime.Seconds() >= float64(seconds) {
			break
		}
	}
}

func main() {
	fmt.Println(1)
	sleep(2*time.Second)
	fmt.Println(2)
	sleep1(2)
	fmt.Println(3)
}