package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, v2 := range nums {
			if i != j && v+v2 == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	//twoSum(nums, target)
	fmt.Println(twoSum(nums, target))
}
