package main

import (
	"fmt"
)

func main() {
	arr := []int{81, 70, 75, 30, 80, 7, 49, 32, 41, 95, 67, 18, 89, 53, 25, 100, 65, 37, 59, 72}

	bubbleSort(arr, func(a, b int) bool {
		return a < b
	})

	fmt.Println(arr)
}

func bubbleSort(arr []int, c func(int, int) bool) {
	sorting := true
	for sorting {
		sorting = false
		for i := 1; i < len(arr); i++ {
			if c(arr[i-1], arr[i]) {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				sorting = true
			}
		}
	}
}
