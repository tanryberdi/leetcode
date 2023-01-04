package main

import "fmt"

func firstMissingPositive(nums []int) int {
	visited := make([]int, len(nums)+1)

	for _, number := range nums {
		if number < 0 || number > len(nums) {
			continue
		}
		visited[number]++
	}

	for i := 1; i <= len(nums); i++ {
		if visited[i] == 0 {
			return i
		}
	}
	return len(nums) + 1
}

func main() {
	numbers := []int{1}
	fmt.Println(firstMissingPositive(numbers))
}
