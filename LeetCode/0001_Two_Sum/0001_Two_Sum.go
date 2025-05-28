import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	nums_map := make(map[int]int)
	for idx1, num := range nums {
		if idx2, ok := nums_map[target-num]; ok {
			return []int{idx2, idx1}
		} else {
			nums_map[num] = idx1
		}
	}
	return nil
}