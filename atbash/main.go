package main

import (
	"fmt"
)

func main() {
	arr := "abcdefghijklmnopqrstuvwxyzZYXWVUTSRQPONMLKJIHGFEDCBA"

	fmt.Println(atbash(arr))
}

func atbash(str string) string {
	retStr := ""

	for _, v := range str {
		if v >= 97 && v <= 122 {
			retStr += string(26 - (v - 97) + 96)
		} else if v >= 65 && v <= 90 {
			retStr += string(26 - (v - 65) + 64)
		}
	}

	return retStr
}
