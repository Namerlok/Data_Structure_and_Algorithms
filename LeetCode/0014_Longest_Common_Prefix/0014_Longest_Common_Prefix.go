package main

import "fmt"

// Start of code to send //

func longestCommonPrefix(strs []string) string {
	result := strs[0]
	for ind := 1; ind < len(strs); ind++ {
		if len(result) > len(strs[ind]) {
			result = result[:len(strs[ind])]
		}
		for subInd := 0; subInd < len(strs[ind]); subInd++ {
			if subInd >= len(result) {
				break
			}

			if strs[ind][subInd] != result[subInd] {
				result = result[:subInd]
				break
			}
		}
		if len(result) == 0 {
			return result
		}
	}
	return result
}

// End of code to send //

func main() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Printf("#1 test: ")
	if result == "fl" {
		fmt.Println("pass")
	} else {
		fmt.Println("failed. Output: \"%s\", Correct: \"fl\"", result)
	}

	strs = []string{"dog", "racecar", "car"}
	result = longestCommonPrefix(strs)
	fmt.Printf("#1 test: ")
	if result == "" {
		fmt.Println("pass")
	} else {
		fmt.Println("failed. Output: \"%s\", Correct: \"\"", result)
	}
}
