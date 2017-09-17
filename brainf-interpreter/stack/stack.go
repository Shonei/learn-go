package stack

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	result := (*s)[len(*s)-1]
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
	return result
}
