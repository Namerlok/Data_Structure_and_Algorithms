package main

import (
	"fmt"
	"strconv"
)

func romanToInt(s string) int {
	var result int
	var diffI bool
	var diffX bool
	var diffC bool
	for _, ch := range s {
		switch ch {
		case 'I':
			result += 1
			diffI = true
		case 'V':
			if diffI {
				result += 3
			} else {
				result += 5
			}
			diffI = false
		case 'X':
			if diffI {
				result += 8
			} else {
				result += 10
			}
			diffI = false
			diffX = true
		case 'L':
			if diffX {
				result += 30
			} else {
				result += 50
			}
			diffX = false
		case 'C':
			if diffX {
				result += 80
			} else {
				result += 100
			}
			diffX = false
			diffC = true
		case 'D':
			if diffC {
				result += 300
			} else {
				result += 500
			}
			diffC = false
		case 'M':
			if diffC {
				result += 800
			} else {
				result += 1000
			}
			diffC = false
		default:
			return -1
		}
	}
	return result
}

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
