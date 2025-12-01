package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func lengthOfLastWord(s string) int {
	var endWorld int

	for len(s) > 0 {
		r, size := utf8.DecodeLastRuneInString(s)

		isSpace := unicode.IsSpace(r)
		if !isSpace && endWorld == 0 {
			endWorld = len(s)
		} else if isSpace && endWorld != 0 {
			break
		}
		s = s[:len(s)-size]
	}
	return endWorld - len(s)
}

func main() {
	s := "Hello World"
	result := lengthOfLastWord(s)
	if result == 5 {
		fmt.Println(s + ": ok")
	} else {
		fmt.Println(s+": fault \tresult =", result)
	}

	s = "   fly me   to   the moon  "
	result = lengthOfLastWord(s)
	if result == 4 {
		fmt.Println(s + ": ok")
	} else {
		fmt.Println(s+": fault \tresult =", result)
	}

	s = "luffy is still joyboy"
	result = lengthOfLastWord(s)
	if result == 6 {
		fmt.Println(s + ": ok")
	} else {
		fmt.Println(s+": fault \tresult =", result)
	}

	s = "a"
	result = lengthOfLastWord(s)
	if result == 1 {
		fmt.Println(s + ": ok")
	} else {
		fmt.Println(s+": fault \tresult =", result)
	}
}
