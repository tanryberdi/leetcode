package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var quadruplets [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			left, right := j+1, len(nums)-1
			for left < right {
				if left > j+1 && nums[left-1] == nums[left] {
					left++
					continue
				}
				if nums[i]+nums[j]+nums[left]+nums[right] == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return quadruplets
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}
