package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	var dfs func(int, []int)
	dfs = func(idx int, temp []int) {
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)

		for i := idx; i < len(nums); i++ {
			if i == idx || nums[i] != nums[i-1] {
				dfs(i+1, append(temp, nums[i]))
			}
		}
	}

	temp := make([]int, 0, len(nums))
	dfs(0, temp)

	return res
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
