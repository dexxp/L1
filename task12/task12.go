package main


import "fmt"


func main() {
	// слайс строк
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// множество
	set := make(map[string]bool)

	for _, v := range arr {
		// делаем строку из слайса ключем и устанавливаем его в true
		set[v] = true
	}

	// проходимся по ключам множества
	for key := range set {
		// выводим элементы множества
		fmt.Println(key)
	}
}