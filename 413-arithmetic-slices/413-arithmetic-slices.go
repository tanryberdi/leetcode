package main

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	sum := 0
	for i := 0; i < len(dp); i++ {
		sum += dp[i]
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println("Output: ", numberOfArithmeticSlices(nums))
}
