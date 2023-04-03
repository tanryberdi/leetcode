package main

import "fmt"

func searchRange(nums []int, target int) []int {
	start := -1
	end := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			start = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			end = i
			break
		}
	}
	return []int{start, end}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	res := searchRange(nums, target)
	fmt.Println(res)
}
