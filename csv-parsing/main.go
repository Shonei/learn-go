package main

import (
	"fmt"
	// "strings"
	"regexp"
)

func main () {
	str := `1,"Que?","Kay?",2,"Si.","Sea? Kay, sea?","No, no, no. Que... ‘what’.",234,"Kay Watt?","Si, que ‘what’.","C.K. Watt?",3,"Yes!","comma,comma, comma , :)"`

	re := regexp.MustCompile("(\",\")|(,\")|(\",)")
	split := re.Split(str, -1)
	for _, v := range split {
		fmt.Println(v)
	}
}