package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 2}

	naive(arr)
}

func naive(arr []int) {
	for i, v := range arr {
		for j := i + 1; j < len(arr); j++ {
			if v == arr[j] {
				fmt.Println(v)
			}
		}
	}
}
