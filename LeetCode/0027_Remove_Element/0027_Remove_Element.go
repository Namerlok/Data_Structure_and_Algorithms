package main

import "fmt"

// Start of code to send //

func removeElement(nums []int, val int) int {
	var posRemove int
	for ind := 0; ind < len(nums); ind++ {
		if nums[ind] != val {
			nums[posRemove] = nums[ind]
			posRemove++
		}
	}
	return posRemove
}

// End of code to send //

func main() {

	nums := []int{3, 2, 2, 3}
	answer := 2
	result := removeElement(nums, 3)
	fmt.Printf("#test case 1: ")
	if result == answer {
		fmt.Println("pass")
	} else {
		fmt.Printf("failed, input: %d, answer: %d, obtained: %d\n", nums, answer, result)
	}

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	answer = 5
	result = removeElement(nums, 2)
	fmt.Printf("#test case 1: ")
	if result == answer {
		fmt.Println("pass")
	} else {
		fmt.Printf("failed, input: %d, answer: %d obtained: %d\n", nums, answer, result)
	}

}
