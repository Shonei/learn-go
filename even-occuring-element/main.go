package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 32, 2, 634, 664, 2, 32}
	r := getEvelElem(arr)
	fmt.Println(r)
}

func getEvelElem(arr []int) int {
	m := map[int]int{}
	result := -1
	for _, v := range arr {
		if m[v] == 0 {
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
	}

	for i, v := range m {
		if v%2 == 0 {
			result = i
		}
	}
	return result
}
