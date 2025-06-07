package main

import (
	"fmt"
	"strconv"
)

// Start of code to send //

func NumberIdentification(ch rune) int {
	switch ch {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return -1
	}
}

func romanToInt(s string) int {
	var result int
	var prev int
	for ind, ch := range s {
		now := NumberIdentification(ch)
		if ind != 0 && prev < now {
			result = result - 2*prev + now
		} else {
			result += now
		}
		prev = now
	}
	return result
}

// End of code to send //

func main() {
	s := "III"
	fmt.Print(s + "== 3: ")
	if romanToInt(s) == 3 {
		fmt.Println("Ok")
	} else {
		fmt.Println("No")
		fmt.Println("retrun value: " + strconv.Itoa(romanToInt(s)))
	}

	s = "LVIII"
	fmt.Print(s + "== 58: ")
	if romanToInt(s) == 58 {
		fmt.Println("Ok")
	} else {
		fmt.Println("No")
		fmt.Println("retrun value: " + strconv.Itoa(romanToInt(s)))
	}

	s = "MCMXCIV"
	fmt.Print(s + "== 1994: ")
	if romanToInt(s) == 1994 {
		fmt.Println("Ok")
	} else {
		fmt.Println("No")
		fmt.Println("retrun value: " + strconv.Itoa(romanToInt(s)))
	}
}
