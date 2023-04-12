package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	sort.Ints(nums)
	generatePermuteUnique(nums, 0, p, &res, used)
	return res
}

func generatePermuteUnique(nums []int, index int, p []int, res *[][]int, used []bool) {
	if index == len(nums) {
		*res = append(*res, append([]int{}, p...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}
		p = append(p, nums[i])
		used[i] = true
		generatePermuteUnique(nums, index+1, p, res, used)
		p = p[:len(p)-1]
		used[i] = false
	}
}

func main() {
	nums := []int{3, 3, 0, 3}
	fmt.Println(permuteUnique(nums))
}
