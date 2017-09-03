package main 

import (
	"fmt"
)

// I was not able to implement it on my own7
// I have failed 
// The solution I found online
// http://www.ardendertat.com/2011/10/18/programming-interview-questions-9-convert-array/

func getIndex(currentIndex, N int)(int) {
	return (currentIndex%3)*N + (currentIndex/3)
}

func main() {
	arr := []string{"a1", "a2", "a3", "a4", "b1", "b2", "b3", "b4", "c1", "c2", "c3", "c4"}
	N := len(arr)/3

    for currentIndex := range arr {
		swapIndex := getIndex(currentIndex, N)
        for swapIndex < currentIndex {
			swapIndex=getIndex(swapIndex, N)
		}
        arr[currentIndex], arr[swapIndex] = arr[swapIndex], arr[currentIndex]
	}

	fmt.Println(arr)
}