package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	start, preVal := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != preVal {
			start, preVal, nums[start] = start+1, nums[i], nums[i]
		}
	}

	return start
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(removeDuplicates(nums))
}
