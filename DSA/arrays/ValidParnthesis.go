package main

import "fmt"

func isValid(s string) bool {
	
	if len(s)%2 != 0 {
		return false
	}

	
	stack := []rune{}
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (ch == ')' && top == '(') || (ch == '}' && top == '{') || (ch == ']' && top == '[') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	// Test cases
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}
