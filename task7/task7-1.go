package main

import (
	"sync"
	"fmt"
)

type Map struct {
	m map[int]int
	mu sync.Mutex
	wg sync.WaitGroup
}


func (m *Map) RecordTenNumbers() {

	for i := 0; i < 10; i++ {
		m.wg.Add(1)
		k := i
		go func() {
			defer m.wg.Done()
			m.mu.Lock()
			m.m[k] = k
			m.mu.Unlock()
		}()	
	}

	m.wg.Wait()
}

func (m *Map) GetMap() map[int]int {
	return m.m
}

func main() {

	m := Map{m: make(map[int]int)}
	
	m.RecordTenNumbers()

	fmt.Println(m.GetMap())
}
