package main

import "fmt"

func moveZeroes(nums []int) {
	i := 0
	for _, val := range nums {
		if val != 0 {
			nums[i] = val
			i++
		}
	}

	for j := i; j < len(nums); j++ {
		nums[j] = 0
	}

	fmt.Println(nums)
}

func main() {
	nums := []int{0}
	moveZeroes(nums)
}
