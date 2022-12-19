package main

import "fmt"

func findDuplicate(nums []int) int {
	// We will store all numbers
	numbers := make(map[int]int)
	for _, number := range nums {
		if _, ok := numbers[number]; ok {
			return number
		} else {
			numbers[number] = 1
		}
	}
	return -1
}

func main() {
	nums := []int{2, 2, 2, 2, 2}
	fmt.Println(findDuplicate(nums))
}
