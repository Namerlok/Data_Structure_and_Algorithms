package main

import (
	"fmt"
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
	usedLetter := make(map[rune]int)
	var posL, maxLen int
	for posR, r := range str {

		if pos, ok := usedLetter[r]; ok {
			for _, val := range str[posL : pos+1] {
				delete(usedLetter, val)
			}
			posL = pos + 1
		}

		usedLetter[r] = posR

		maxLen = max(maxLen, posR-posL+1)

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
	fmt.Printf("#test case 4: ")
	if result == answer {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("failed, input string: %s, answer: %d, give result: %d\n", str, answer, result)
	}

	str = "abba"
	answer = 2
	result = lengthOfLongestSubstring(str)
	fmt.Printf("#test case 5: ")
	if result == answer {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("failed, input string: %s, answer: %d, give result: %d\n", str, answer, result)
	}
}
