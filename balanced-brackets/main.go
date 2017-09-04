package main

func main() {
	isBalanced("()[]{}(([])){[()][]}")
}

func isBalanced(str string) bool {
	stack := []string{}
	bMap := map[string]string{
		"[": "]",
		"]": "[",
		"{": "}",
		"}": "{",
		"(": ")",
		")": "("}

	for i := 0; i < len(str); i++ {
		if len(stack) == 0 {
			stack = append(stack, string(str[i]))
		} else if bMap[string(str[i])] == string(stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(str[i]))
		}
	}

	return len(stack) == 0
}
