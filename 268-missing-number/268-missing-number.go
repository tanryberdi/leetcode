package main

import "fmt"

func missingNumber(nums []int) int {
	// sum is sum of all number 0..n
	sum := (len(nums) * (len(nums) + 1)) / 2
	for _, number := range nums {
		sum -= number
	}
	return sum
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}
