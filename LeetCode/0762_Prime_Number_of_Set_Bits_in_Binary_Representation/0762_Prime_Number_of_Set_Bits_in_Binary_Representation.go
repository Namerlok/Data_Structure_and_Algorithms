package main

import "math"

func isPrimeNum(num int) bool {
	if num < 2 {
		return false
	}
	for it := 2; it <= int(math.Sqrt(float64(num))); it++ {
		if num%it == 0 {
			return false
		}
	}
	return true
}

func countUnitInNum(num int) int {
	var result int
	for num > 0 {
		if num&1 == 1 {
			result++
		}
		num >>= 1
	}
	return result
}

func countPrimeSetBits(left int, right int) int {
	var result int
	for it := left; it <= right; it++ {
		if isPrimeNum(countUnitInNum(it)) {
			result++
		}
	}
	return result
}
