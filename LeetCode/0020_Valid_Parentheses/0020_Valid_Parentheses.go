package main

import "fmt"

// Start of code to send //

type node struct {
	ch   rune
	next *node
}

type stack struct {
	head *node
	len  int
}

func (st *stack) Push(newRune rune) {
	*st = stack{&node{newRune, st.head}, st.len + 1}
}

func (st *stack) Pop() (res rune) {
	if st.head == nil {
		return '\x00'
	}
	res = st.head.ch
	st.len--
	st.head = st.head.next
	return res
}

func isValid(str string) bool {
	var st stack
	for _, ch := range str {
		switch ch {
		case '(':
			st.Push(')')
		case ')':
			if st.Pop() != ')' {
				return false
			}
		case '{':
			st.Push('}')
		case '}':
			if st.Pop() != '}' {
				return false
			}
		case '[':
			st.Push(']')
		case ']':
			if st.Pop() != ']' {
				return false
			}
		default:
		}
	}
	if st.len == 0 {
		return true
	} else {
		return false
	}
}

// End of code to send //

func main() {

	str := "()"
	answer := true
	fmt.Printf("#test case 1: ")
	if isValid(str) == answer {
		fmt.Println("pass")
	} else {
		fmt.Printf("failed, input: %s, answer: %t, obtained: %t\n", str, answer, isValid(str))
	}

	str = "()[]{}"
	answer = true
	fmt.Printf("#test case 2: ")
	if isValid(str) == answer {
		fmt.Println("pass")
	} else {
		fmt.Printf("failed, input: %s, answer: %t, obtained: %t\n", str, answer, isValid(str))
	}

	str = "(]"
	answer = false
	fmt.Printf("#test case 3: ")
	if isValid(str) == answer {
		fmt.Println("pass")
	} else {
		fmt.Printf("failed, input: %s, answer: %t, obtained: %t\n", str, answer, isValid(str))
	}
}
