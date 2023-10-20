package main

import (
	"sync"
	"fmt"
)

// стурктура Map 
type Map struct {
	m map[int]int // общий ресурс
	mu sync.Mutex  // мьютекс
	wg sync.WaitGroup 
}

// метод записывающий 10 значений в мапу
func (m *Map) RecordTenNumbers() {

	// добавляем один к счетчику запущенных горутин
	for i := 0; i < 10; i++ {
		m.wg.Add(1)
		k := i
		go func() {
			// отнимаем один от счетчика после завершения горутины
			defer m.wg.Done()
			m.mu.Lock() // блокируем доступ к мапе
			m.m[k] = k // записываем значение
			m.mu.Unlock() // деблокируем доступ
		}()	
	}
	// дожидаемся выполнения всех горутин
	m.wg.Wait()
}

// метод возвращающий мапу структуры Map
func (m *Map) GetMap() map[int]int {
	return m.m
}

func main() {
	// создаём экземпляр Map
	m := Map{m: make(map[int]int)}
	
	// конкурентно записываем 10 значений в мапу
	m.RecordTenNumbers()

	// выводим мапу
	fmt.Println(m.GetMap())
}
