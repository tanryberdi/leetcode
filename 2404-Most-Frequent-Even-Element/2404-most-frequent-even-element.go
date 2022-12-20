package main

import "fmt"

func mostFrequentEven(nums []int) int {
	numbers := make(map[int]int)

	for _, number := range nums {
		numbers[number]++
	}

	resultKey := 0
	resultValue := 0
	found := false
	for key, value := range numbers {
		if key%2 == 0 && value >= resultValue {
			found = true
			if value == resultValue && key > resultKey {
				continue
			}
			resultKey = key
			resultValue = value
		}
	}

	if !found {
		return -1
	}

	return resultKey
}

func main() {
	numbers := []int{0, 0, 0, 0, 0}
	fmt.Println(mostFrequentEven(numbers))
}
