package main

import (
	"fmt"
)

func threeSumClosest(nums []int, target int) int {
	min := absoluteValue(target, nums[0]+nums[1]+nums[2])
	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if absoluteValue(target, nums[i]+nums[j]+nums[k]) < min {
					min = absoluteValue(target, nums[i]+nums[j]+nums[k])
					result = nums[i] + nums[j] + nums[k]
					fmt.Println(i, " ", j, " ", k)
				}
			}
		}
	}

	return result
}

func absoluteValue(a, b int) int {
	if a-b < 0 {
		return -1 * (a - b)
	}

	return a - b
}

func main() {
	numbers := []int{0, 0, 0}
	target := 1
	fmt.Println(threeSumClosest(numbers, target))
}
