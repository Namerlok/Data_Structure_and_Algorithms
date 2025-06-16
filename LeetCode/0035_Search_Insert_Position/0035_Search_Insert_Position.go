package main

import "fmt"

// Start of code to send //

func searchInsert(nums []int, target int) int {
	posL, posR := 0, len(nums)-1
	for posR-posL > 1 {
		privot := (posR + posL) / 2
		fmt.Printf("privot: %d, posL: %d, posR:%d \n", privot, posL, posR)
		if target == nums[privot] {
			return privot
		} else if target > nums[privot] {
			posL = privot
		} else {
			posR = privot
		}
	}

	if target <= nums[posL] {
		return posL
	} else if nums[posL] < target && target <= nums[posR] {
		return posR
	} else {
		return posR + 1
	}

}

// End of code to send //

func main() {
	nums := []int{1, 3, 5, 6}
	answer := 2
	result := searchInsert(nums, 5)
	fmt.Printf("#test case 1: ")
	if result == answer {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("failed, input string: %d, answer: %d, give result: %d\n", nums, answer, result)
	}

	nums = []int{1, 3, 5, 6}
	answer = 1
	result = searchInsert(nums, 2)
	fmt.Printf("#test case 1: ")
	if result == answer {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("failed, input string: %d, answer: %d, give result: %d\n", nums, answer, result)
	}

	nums = []int{1, 3, 5, 6}
	answer = 4
	result = searchInsert(nums, 7)
	fmt.Printf("#test case 1: ")
	if result == answer {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("failed, input string: %d, answer: %d, give result: %d\n", nums, answer, result)
	}

	nums = []int{1}
	answer = 0
	result = searchInsert(nums, 1)
	fmt.Printf("#test case 1: ")
	if result == answer {
		fmt.Printf("pass\n")
	} else {
		fmt.Printf("failed, input string: %d, answer: %d, give result: %d\n", nums, answer, result)
	}
}
