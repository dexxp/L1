package main

import "fmt"

func quickSort(arr []int) []int {
    // если массив из одного элемента или меньше, возвращаем его
    if len(arr) <= 1 {
        return arr
    }

    // в качестве опорного элемента взяли последний элемент
    pivot := arr[len(arr)-1]
    // слайс для чисел меньших опорного элемента
    left := make([]int, 0)
    // слайс для чисел больших опорного элемента
    right := make([]int, 0)
    // слайс для чисел равных опорному элемента
    equal := make([]int, 0)

    // проходимся по слайсу
    for _, num := range arr {
        // если меньше опорного элемента
        if num < pivot {
            // добавляем в слайс left
            left = append(left, num)
        } else if num > pivot {
            right = append(right, num)
        } else {
            equal = append(equal, num)
        }
    }

    // вызываем рекурсивно функцию для слайсов left и right
    left = quickSort(left)
    right = quickSort(right)

    // склеиваем полученнае результаты в один массив
    return append(append(left, equal...), right...)
}

func main() {
    // слайс чисел
    arr := []int{7, 2, 1, 6, 8, 5, 3, 4}
    // сортируем слайс
    sortedArr := quickSort(arr)
    // выводим отсортированный массив
    fmt.Println(sortedArr)
}