package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 6, 7, 8}
	c := []int{}
	

	for _, nA := range a {
		for _, nB := range b {
			if nA == nB {
				found := false
				for _, nC := range c {
					if nA == nC {
						found = true
						break
					}
				}
				if !found {
					c = append(c, nA)
				}
			}
		}
	}

	fmt.Println(c)

}