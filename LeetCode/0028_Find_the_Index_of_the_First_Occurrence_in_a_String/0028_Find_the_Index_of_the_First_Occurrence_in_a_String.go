package main

import (
	"fmt"
)

// Start of code to send //

func PrefixFunction(needle []rune) []int {
	pi := make([]int, len(needle))
	for pos := 1; pos < len(needle); pos++ {
		ind := pi[pos-1]
		for ind > 0 && needle[ind] != needle[pos] {
			ind = pi[ind-1]
		}
		if needle[ind] == needle[pos] {
			ind++
		}
		pi[pos] = ind
	}
	return pi
}

// Knuth-Morris-Pratt algorithm
func strStr(haystack string, needle string) int {
	text := []rune(haystack)
	str := []rune(needle)
	prefixFunc := PrefixFunction(str)
	posStr := 0
	for posText := 0; posText < len(text); posText++ {
		for posStr > 0 && str[posStr] != text[posText] {
			posStr = prefixFunc[posStr-1]
		}
		if str[posStr] == text[posText] {
			posStr++
		}
		if posStr == len(needle) {
			return posText - posStr + 1
		}
	}
	return -1
}

// End of code to send //

func main() {

	haystack := "sadbutsad"
	needle := "sad"
	answer := 0
	result := strStr(haystack, needle)
	fmt.Printf("#test case 1: ")
	if result == answer {
		fmt.Println("pass")
	} else {
		fmt.Printf("failed, input: [haystack: %s], [needle: %s], answer: %d, obtained: %d\n",
			haystack, needle, answer, result)
	}

	haystack = "leetcode"
	needle = "leeto"
	answer = -1
	result = strStr(haystack, needle)
	fmt.Printf("#test case 2: ")
	if result == answer {
		fmt.Println("pass")
	} else {
		fmt.Printf("failed, input: [haystack: %s], [needle: %s], answer: %d, obtained: %d\n",
			haystack, needle, answer, result)
	}

	haystack = "hello"
	needle = "ll"
	answer = 2
	result = strStr(haystack, needle)
	fmt.Printf("#test case 3: ")
	if result == answer {
		fmt.Println("pass")
	} else {
		fmt.Printf("failed, input: [haystack: %s], [needle: %s], answer: %d, obtained: %d\n",
			haystack, needle, answer, result)
	}

	haystack = "mississippi"
	needle = "issip"
	answer = 4
	result = strStr(haystack, needle)
	fmt.Printf("#test case 4: ")
	if result == answer {
		fmt.Println("pass")
	} else {
		fmt.Printf("failed, input: [haystack: %s], [needle: %s], answer: %d, obtained: %d\n",
			haystack, needle, answer, result)
	}
}
