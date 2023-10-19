package main

import (
	"fmt"
	"math/rand"
)

// Примерная реализация функции createHugeString
func createHugeString(size int) string {
	buf := make([]byte, size)

	for i := 0; i < size; i++ {
		buf[i] = byte(rand.Intn(256))
	}

	return string(buf)
}

func someFunc() {
	// вызываем функцию, которая возвращает большую строку и эта строка может занимать много памяти
	v := createHugeString(1 << 10)
	// проверяем, что мы не выходим за рамки строки
	if len(v) >= 99 {
		// забираем только первые 100 символов из большой строки
		justString := v[:100]
		fmt.Println(justString)
	}
	// очищаем память, которую занимает строка
	v = ""
}

func main() {
	someFunc()
}