package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k; j++ {
			if len(nums) > j && nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	numbers := []int{99, 99}
	k := 2
	fmt.Println(containsNearbyDuplicate(numbers, k))
}
