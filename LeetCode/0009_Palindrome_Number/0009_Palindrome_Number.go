package main

import (
	"fmt"
	"math"
)

// Start of code to send //

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	posLeft := 1
	for i := 0; i < int(math.Log10(float64(x))); i++ {
		posLeft *= 10
	}
	posRight := 1

	for posLeft > posRight {
		if (x/posLeft)%10 != (x/posRight)%10 {
			return false
		} else {
			posLeft /= 10
			posRight *= 10
		}
	}
	return true
}

// End of code to send //

func main() {
	x := 121
	fmt.Printf("%d is Palindrome: %t\n", x, isPalindrome(x))
	x = -121
	fmt.Printf("%d is Palindrome: %t\n", x, isPalindrome(x))
	x = 10
	fmt.Printf("%d is Palindrome: %t\n", x, isPalindrome(x))
}
