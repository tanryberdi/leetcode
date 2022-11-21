package main

import "fmt"

func singleNumber(nums []int) int {
	numbers := make(map[int]int)

	for _, num := range nums {
		if _, ok := numbers[num]; ok {
			delete(numbers, num)
		} else {
			numbers[num] = 1
		}
	}

	var key int
	for key = range numbers {
		break
	}

	return key
}

func main() {
	nums := []int{1}
	fmt.Println(singleNumber(nums))
}
