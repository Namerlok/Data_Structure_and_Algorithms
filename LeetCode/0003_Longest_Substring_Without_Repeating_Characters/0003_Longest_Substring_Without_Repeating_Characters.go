package main

import (
	"fmt"
	"unicode/utf8"
)

// Start of code to send //

func max(left, right int) int {
	if left > right {
		return left
	} else {
		return right
	}
}

func lengthOfLongestSubstring(str string) int {
	usedLetter := make(map[rune]struct{})
	var posL, maxLen, len int
	for _, r := range str {
		if _, ok := usedLetter[r]; ok {
			for {
				val, size := utf8.DecodeRuneInString(str[posL:])
				fmt.Printf("val: %c\n", val)
				delete(usedLetter, val)
				posL += size
				len--
				if val == r {
					break
				}
			}
		}
		usedLetter[r] = struct{}{}
		len++

		maxLen = max(maxLen, len)

		fmt.Printf("map: ")
		for key, _ := range usedLetter {
			fmt.Printf("%c ", key)
		}
		fmt.Printf("\n")
	}
	return maxLen
}

// End of code to send //

func main() {
	str := "abcabcbb"
	answer := 3
	result := lengthOfLongestSubstring(str)
	fmt.Printf("#test case 1: ")
	if result == answer {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("failed, input string: %s, answer: %d, give result: %d\n", str, answer, result)
	}

	str = "bbbbb"
	answer = 1
	result = lengthOfLongestSubstring(str)
	fmt.Printf("#test case 2: ")
	if result == answer {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("failed, input string: %s, answer: %d, give result: %d\n", str, answer, result)
	}

	str = "pwwkew"
	answer = 3
	result = lengthOfLongestSubstring(str)
	fmt.Printf("#test case 3: ")
	if result == answer {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("failed, input string: %s, answer: %d, give result: %d\n", str, answer, result)
	}

	str = "aab"
	answer = 2
	result = lengthOfLongestSubstring(str)
	fmt.Printf("#test case 34: ")
	if result == answer {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("failed, input string: %s, answer: %d, give result: %d\n", str, answer, result)
	}
}
