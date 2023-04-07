package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	sum := 0
	ans := len(nums) + 1
	for right, v := range nums {
		sum += v
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(target, nums))
}
