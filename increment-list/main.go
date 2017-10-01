package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 9, 9, 9}
	arr2 := []int{7, 2, 6, 4}
	arr3 := []int{7, 2, 6, 9}
	increment(arr)
	increment(arr2)
	increment(arr3)
	fmt.Println(arr, arr2, arr3)
}

func increment(arr []int) {
	arr[len(arr)-1]++

	for i := len(arr) - 1; i >= 0 && arr[i] >= 10; i-- {
		if i == 0 {
			arr = append([]int{1}, arr...)
		} else {
			arr[i] = 0
			arr[i-1]++
		}
	}
}
