package main

import "fmt"

func main() {
	arr := []int{43, 76, 5, 34, 5, 4, 6, 325, 4, 65, 345, 5, 3, 5, 6, 4, 35, 78, 574, 568, 15, 35, 7465, 843, 524, 64, 73}

	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quicksort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quicksort(arr, low, p-1)
		quicksort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	if arr[high] < arr[i+1] {
		arr[i+1], arr[high] = arr[high], arr[i+1]
	}

	return i + 1
}
