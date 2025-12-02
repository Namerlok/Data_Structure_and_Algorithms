package main

import (
	"fmt"
	"unicode/utf8"
)

func convert(s string, numRows int) string {
	result := make([]rune, len(s))
	pos := 0
	iter := 0
	count := 0

	for id := 0; id < len(result); id++ {
		if pos >= len(result) {
			iter++
			pos = iter
			count = 0
		}

		diff := 0
		if count%2 == 0 && iter != numRows-1 || iter == 0 {
			diff = (numRows - iter - 1) * 2
		} else {
			diff = iter * 2
		}
		count++

		if diff == 0 {
			diff = 1
		}

		// fmt.Printf("\n pos = %d, iter = %d, diff = %d \n", pos, iter, diff)
		result[id], _ = utf8.DecodeRuneInString(s[pos:])
		pos += diff
	}
	return string(result)
}

func main() {
	inputStr := "PAYPALISHIRING"
	numRows := 3
	expectedResult := "PAHNAPLSIIGYIR"
	fmt.Printf("input string \"%s\" result is ", inputStr)
	result := convert(inputStr, numRows)
	if result == expectedResult {
		fmt.Println("OK")
	} else {
		fmt.Println("fault. Received \"", result,
			"\", but expected \"", expectedResult, "\"")
	}

	inputStr = "PAYPALISHIRING"
	numRows = 4
	expectedResult = "PINALSIGYAHRPI"
	fmt.Printf("input string \"%s\" result is ", inputStr)
	result = convert(inputStr, numRows)
	if result == expectedResult {
		fmt.Println("OK")
	} else {
		fmt.Println("fault. Received \"", result,
			"\", but expected \"", expectedResult, "\"")
	}

	inputStr = "A"
	numRows = 1
	expectedResult = "A"
	fmt.Printf("input string \"%s\" result is ", inputStr)
	result = convert(inputStr, numRows)
	if result == expectedResult {
		fmt.Println("OK")
	} else {
		fmt.Println("fault. Received \"", result,
			"\", but expected \"", expectedResult, "\"")
	}

	inputStr = "AB"
	numRows = 1
	expectedResult = "AB"
	fmt.Printf("input string \"%s\" result is ", inputStr)
	result = convert(inputStr, numRows)
	if result == expectedResult {
		fmt.Println("OK")
	} else {
		fmt.Println("fault. Received \"", result,
			"\", but expected \"", expectedResult, "\"")
	}
}
