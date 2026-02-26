package main

func plusOne(digits []int) []int {
	one := 1
	for ind := len(digits) - 1; one != 0 && ind >= 0; ind-- {
		if digits[ind] == 9 {
			digits[ind] = 0
		} else {
			digits[ind] += one
			one = 0
		}
	}

	if one != 0 {
		result := make([]int, len(digits)+1)
		result[0] = 1
		return result
	} else {
		return digits
	}
}
