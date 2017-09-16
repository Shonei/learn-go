package inter

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Shonei/learn-go/brainf-interpreter/stack"
)

type File struct {
	str    []byte
	index  int
	memory []byte
	stack  stack.Stack
}

func (f *File) ReadFile(s string) {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	f.str = data
}

func (f *File) Interpret() {
	f.memory = append(f.memory, 0)
	for i := 0; i < len(f.str); i++ {
		// fmt.Println(f.str[i])
		switch f.str[i] {
		case 60: // <
			f.index--
		case 62: // >
			f.index++
			if len(f.memory) >= f.index {
				f.memory = append(f.memory, 0)
			}
		case 43: // +
			f.memory[f.index]++
		case 45: // -
			f.memory[f.index]--
		case 46: // .
			fmt.Printf("%s ", string(f.memory[f.index]))
		case 44: // ,
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			f.memory[f.index] = []byte(input)[0]
		case 91: // [
			if f.memory[f.index] == 0 {
				for ; f.str[i] != 93; i++ {
				}
			} else {
				f.stack.Push(i)
				// i++
			}
		case 93: // ]
			// fmt.Printf("%v", f.memory[f.index])
			if f.memory[f.index] == 0 {
				f.stack.Pop()
				// i++
			} else {
				i = f.stack.Top()
			}
		default:
			fmt.Printf("")
		}
	}
}
