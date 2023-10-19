package main

import "fmt"

func quickSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[len(arr)-1]
    left := make([]int, 0)
    right := make([]int, 0)
    equal := make([]int, 0)

    for _, num := range arr {
        if num < pivot {
            left = append(left, num)
        } else if num > pivot {
            right = append(right, num)
        } else {
            equal = append(equal, num)
        }
    }

    left = quickSort(left)
    right = quickSort(right)

    return append(append(left, equal...), right...)
}

func main() {
    arr := []int{7, 2, 1, 6, 8, 5, 3, 4}
    sortedArr := quickSort(arr)
    fmt.Println(sortedArr)
}