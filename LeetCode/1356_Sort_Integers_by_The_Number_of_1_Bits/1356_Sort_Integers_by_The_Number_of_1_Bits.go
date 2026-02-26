package main

import "sort"

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(left, right int) bool {
		left = arr[left]
		right = arr[right]
		var isLess bool
		if left < right {
			isLess = true
		}
		var countLeft, countRight int
		for left > 0 || right > 0 {
			if left > 0 && left&1 == 1 {
				countLeft++
			}
			if right > 0 && right&1 == 1 {
				countRight++
			}
			left >>= 1
			right >>= 1
		}
		if countLeft < countRight {
			return true
		} else if countLeft > countRight {
			return false
		} else { // if countLeft == countRight
			return isLess
		}
	})
	return arr

}
