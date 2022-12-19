package main

import "fmt"

func majorityElement(nums []int) int {
	numbers := make(map[int]int)
	max := 1
	value := 0
	for _, number := range nums {
		numbers[number]++
		if max <= numbers[number] {
			max = numbers[number]
			value = number
		}
	}
	return value
}

func main() {
	nums := []int{1}
	fmt.Println(majorityElement(nums))
}
