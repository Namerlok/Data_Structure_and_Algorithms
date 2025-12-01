package main

import "fmt"

func longestPalindrome(s string) string {
	maxLenPal := 1
	begintPal := 0
	for i := 0; i < len(s)-(maxLenPal-1)/2; i++ {
		expand := func(l, r int) {
			for ; l > 0 && r < len(s)-1; l, r = l-1, r+1 {
				if s[l-1] != s[r+1] {
					break
				}
			}
			if maxLenPal < r-l+1 {
				maxLenPal = r - l + 1
				begintPal = l
			}
		}

		//случай четного палиндрома
		expand(i, i)

		//случай нечетного палиндрома
		if i+1 < len(s) && s[i] == s[i+1] {
			expand(i, i+1)
		}
	}
	return s[begintPal : begintPal+maxLenPal]
}

func main() {
	s := "babad"
	result := longestPalindrome(s)
	fmt.Printf("result = %s \t", result)
	if result == "bab" || result == "aba" {
		fmt.Println(s + ": ok")
	} else {
		fmt.Println(s + ": fault")
	}

	s = "cbbd"
	result = longestPalindrome(s)
	fmt.Printf("result = %s \t", result)
	if result == "bb" {
		fmt.Println(s + ": ok")
	} else {
		fmt.Println(s + ": fault")
	}

	s = "a"
	result = longestPalindrome(s)
	fmt.Printf("result = %s \t", result)
	if result == "a" {
		fmt.Println(s + ": ok")
	} else {
		fmt.Println(s + ": fault")
	}
	s = "bb"
	result = longestPalindrome(s)
	fmt.Printf("result = %s \t", result)
	if result == "bb" {
		fmt.Println(s + ": ok")
	} else {
		fmt.Println(s + ": fault")
	}

	s = "ccc"
	result = longestPalindrome(s)
	fmt.Printf("result = %s \t", result)
	if result == "ccc" {
		fmt.Println(s + ": ok")
	} else {
		fmt.Println(s + ": fault")
	}

	s = "dddd"
	result = longestPalindrome(s)
	fmt.Printf("result = %s \t", result)
	if result == "dddd" {
		fmt.Println(s + ": ok")
	} else {
		fmt.Println(s + ": fault")
	}
}
