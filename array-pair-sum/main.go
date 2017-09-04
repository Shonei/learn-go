package main

import (
	"fmt"
)

type pair struct {
	i int
	j int
}

func main() {
	arr := []int{3, 4, 5, 4, 4}
	fmt.Println(getSumPairs(arr, 8))
}

func getSumPairs(arr []int, sum int) []pair {
	pairsArr := []pair{}
	lenght := len(arr)
	for i := 0; i < lenght; i++ {
		for j := 1; j < len(arr); j++ {
			if arr[0]+arr[j] == sum {
				pairsArr = append(pairsArr, pair{arr[0], arr[j]})
			}
		}
		arr = arr[1:]
	}
	return pairsArr
}
