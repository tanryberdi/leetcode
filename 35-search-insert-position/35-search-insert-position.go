package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	tail := 0
	head := len(nums) - 1

	if nums[head] < target {
		return head + 1
	}

	if nums[tail] > target {
		return 0
	}

	for {
		middle := (tail + head) / 2
		// target found
		if nums[middle] == target {
			return middle
		}

		if tail > head {
			return head + 1
		}

		if nums[middle] < target {
			tail = middle + 1
		} else {
			head = middle - 1
		}
	}
}

func main() {
	// 1,3,5,6
	nums := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(searchInsert(nums, target))
}
