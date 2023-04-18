package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	count := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
	}
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		maxLen = max(maxLen, dp[i])
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] == maxLen {
			res += count[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findNumberOfLIS(nums))
}
