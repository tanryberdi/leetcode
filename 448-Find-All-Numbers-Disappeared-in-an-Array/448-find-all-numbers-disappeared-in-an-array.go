package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	var result []int
	numbers := make(map[int]int)

	for _, number := range nums {
		numbers[number]++
	}

	for index := 1; index <= len(nums); index++ {
		if _, ok := numbers[index]; !ok {
			result = append(result, index)
		}
	}

	return result
}

func main() {
	numbers := []int{2}
	fmt.Println(findDisappearedNumbers(numbers))
}
