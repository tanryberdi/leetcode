package main

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	max := 0
	step := 0
	end := 0
	for i := 0; i < n-1; i++ {
		if max < nums[i]+i {
			max = nums[i] + i
		}
		if i == end {
			end = max
			step++
		}
	}
	return step
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}
