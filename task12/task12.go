package main


import "fmt"


func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)

	for _, v := range arr {
		set[v] = true
	}

	for key := range set {
		fmt.Println(key)
	}
}