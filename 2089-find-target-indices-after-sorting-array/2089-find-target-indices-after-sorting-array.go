package main

import (
	"fmt"
	"sort"
)

func targetIndices(nums []int, target int) []int {
	var result []int

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 5, 2, 3}
	target := 7
	fmt.Println(targetIndices(nums, target))
}
