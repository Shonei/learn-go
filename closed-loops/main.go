package main

import (
	"fmt"
	"strconv"
)

func getLoopCount(i int32) int {
	if i == 57 ||
		i == 48 ||
		i == 54 {
		return 1
	} else if i == 56 {
		return 2
	}

	return 0
}

func countLoops(number int) int {
	str := strconv.FormatUint(uint64(number), 10)

	sum := 0
	for _, v := range str {
		sum += getLoopCount(v)
	}

	return sum
}

func main() {
	integer := 124346568
	result := countLoops(integer)
	fmt.Println(result)
}
