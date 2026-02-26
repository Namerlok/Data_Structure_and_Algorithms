package main

import "sort"

const MaxBitsInNum = 14

func sortByBits(arr []int) []int {
	countBit := make([][]int, MaxBitsInNum)
	for pos, el := range arr {
		count := 0
		for el > 0 {
			if el&1 == 1 {
				count++
			}
			el >>= 1
		}
		countBit[count] = append(countBit[count], arr[pos])
	}
	pos := 0
	for _, subArr := range countBit {
		sort.Slice(subArr, func(i, j int) bool {
			if subArr[i] < subArr[j] {
				return true
			} else {
				return false
			}
		})
		for _, el := range subArr {
			arr[pos] = el
			pos++
		}

	}
	return arr
}
