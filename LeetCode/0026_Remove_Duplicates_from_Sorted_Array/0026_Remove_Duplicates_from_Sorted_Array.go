package main

import "fmt"

// Start of code to send //

func removeDuplicates(nums []int) int {
	posUnique := 0

	for _, val := range nums {
		if nums[posUnique] != val {
			posUnique++
			nums[posUnique] = val
		}
	}
	return posUnique + 1
}

// End of code to send //

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums), nums)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums), nums)
}
