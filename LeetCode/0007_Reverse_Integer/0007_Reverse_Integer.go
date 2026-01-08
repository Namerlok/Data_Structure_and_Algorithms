package main

import (
	"fmt"
	"math"
)

func reverse(x int32) int32 {
	var result int32
	for x != 0 {
		if (result >= 0 && (result > math.MaxInt32/10 || (result == math.MaxInt32/10 && x%10 > math.MaxInt32%10))) ||
			(result < 0 && (result < math.MinInt32/10 || (result == math.MinInt32/10 && x%10 < math.MinInt32%10))) {
			return 0
		}
		result = result*10 + x%10
		x /= 10
	}
	return result
}

func main() {
	result := reverse(123)
	fmt.Println(-123 / 10)
	fmt.Println(-123 % 10)
	if result != 321 {
		fmt.Println("result is uncorrect: ", result)
	} else {
		fmt.Println("Test \"123\" Ok")
	}
	result = reverse(-123)
	if result != -321 {
		fmt.Println("result is uncorrect: ", result)
	} else {
		fmt.Println("Test \"-123\" Ok")
	}
	result = reverse(120)
	if result != 21 {
		fmt.Println("result is uncorrect: ", result)
	} else {
		fmt.Println("Test \"120\" Ok")
	}
}
