package main

import "github.com/Shonei/learn-go/brainf-interpreter/inter"

func main() {
	inter := inter.File{}
	inter.ReadFile("C:\\Users\\Teodor\\Documents\\go\\src\\github.com\\Shonei\\learn-go\\brainf-interpreter\\helloWorld.bf")
	inter.Interpret()
}
