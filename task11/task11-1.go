package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 6, 7, 8}
	c := []int{}

	c = append(c, a...)
	c = append(c, b...)

	m := make(map[int]int)

	for _, n := range c {
		_, ok := m[n]
		if !ok {
			m[n] = 0
		}

		m[n]++
	}

	for k, v := range m {
		if v >= 2 {
			fmt.Println(k)
		}
	}

}