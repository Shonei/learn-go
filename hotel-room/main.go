package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		arr := getDivisors(i)
		s := sum(arr)

		if s > i {
			if isMyRoom(arr, i, s) {

				fmt.Println(arr, i)
			}
			// fmt.Println(i)
		}
	}
}

func getDivisors(val int) []int {
	result := []int{}

	for i := 1; i <= val/2; i++ {
		if val%i == 0 {
			result = append(result, i)
		}
	}

	return result
}

func sum(val []int) int {
	result := 0
	for _, v := range val {
		result += v
	}

	return result
}

func isMyRoom(arr []int, val, s int) bool {
	result := true

	temp := s
	for i := len(arr) - 1; i >= 0; i-- {
		temp = temp - arr[i]
		if temp < val {
			temp = temp + arr[i]
		} else if temp == val {
			result = false
		}
	}

	return result
}
