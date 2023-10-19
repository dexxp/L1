package main

import "fmt"

func main() {
	tmps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int64][]float64)

	for _, tmp := range tmps {
		k := int64(tmp) / 10 * 10
		_, ok := m[k]
		if !ok {
			m[k] = []float64{}
		}
		fmt.Println(tmp)
		m[k] = append(m[k], tmp)
	}

	fmt.Println(m)
}