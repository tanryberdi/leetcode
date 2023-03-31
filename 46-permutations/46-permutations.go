package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	var helper func([]int, int)
	helper = func(nums []int, start int) {
		if start == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			helper(nums, start+1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	helper(nums, 0)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
