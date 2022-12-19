package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for _, x := range nums {
		if x != val {
			nums[i] = x
			i++
		}
	}
	nums = nums[:i]

	return i
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}
