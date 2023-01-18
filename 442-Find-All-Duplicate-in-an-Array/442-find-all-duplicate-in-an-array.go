package main

import (
	"fmt"
)

func findDuplicates(nums []int) []int {
	var result []int
	numbers := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numbers[nums[i]]++
	}

	for key, value := range numbers {
		if value > 1 {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	numbers := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(numbers))
}
