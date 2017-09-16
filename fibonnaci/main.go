package main

import (
	"fmt"
)

func main() {
	fmt.Println(recursive(6))
	fmt.Println(iterative(6))
}

func recursive(index int) int {
	if index <= 1 {
		return 1
	}

	return recursive(index-1) + recursive(index-2)
}

func iterative(index int) int {
	if index < 2 {
		return -1
	}

	prev1 := 1
	prev2 := 2
	result := 0
	for i := 2; i < index; i++ {
		result = prev1 + prev2
		prev1 = prev2
		prev2 = result
	}

	return result
}
