package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(canScramble("abc", "cba"))
	fmt.Println(canScramble("aab", "bba"))
}

func canScramble(original, scrambled string) bool {
	var result bool
	result = true
	for _, v := range scrambled {
		if strings.ContainsRune(original, v) {
			// prevents duplicating characters
			original = strings.Replace(original, string(v), "", 1)
		} else {
			result = false
			break
		}
	}

	return result
}
