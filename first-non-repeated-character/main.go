package main

import "fmt"

func main() {
	fmt.Println(firstSingleChar("AABBCCEEFFDF"))
}

func firstSingleChar(s string) string {
	if s[0] != s[1] {
		return string(s[0])
	}

	for i := 1; i < len(s)-1; i++ {
		if s[i] != s[i-1] && s[i] != s[i+1] {
			return string(s[i])
		}
	}

	if s[len(s)-1] != s[len(s)-2] {
		return string(s[len(s)-1])
	}
	return ""
}
