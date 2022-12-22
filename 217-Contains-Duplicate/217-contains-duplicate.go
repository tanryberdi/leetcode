package main

import "fmt"

func containsDuplicate(nums []int) bool {
	numbers := make(map[int]int)

	for _, number := range nums {
		numbers[number]++
		if numbers[number] > 1 {
			return true
		}
	}
	return false
}

func main() {
	numbers := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(numbers))
}
