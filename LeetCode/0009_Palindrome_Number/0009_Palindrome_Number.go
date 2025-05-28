package main

import (
	"fmt"
)

// Start of code to send //

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	rev := 0
	for rev < x {
		rev = rev*10 + x%10
		x /= 10
	}

	return (rev == x) || (rev/10 == x)
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
