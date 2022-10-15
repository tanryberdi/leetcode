package main

import "fmt"

func addToArrayForm(num []int, k int) []int {
	var digits []int
	for k > 0 {
		digits = append([]int{k % 10}, digits...)
		k /= 10
	}

	i := len(num) - 1
	j := len(digits) - 1
	step := i
	if j > step {
		step = j
	}
	result := make([]int, step+1)

	if j < 0 {
		return num
	}

	sum := 0
	remainder := 0
	for {
		if i < 0 && j < 0 {
			break
		}

		if j >= 0 && i >= 0 {
			sum = num[i] + digits[j] + remainder
		} else {
			if i < 0 {
				sum = digits[j] + remainder
			} else {
				sum = num[i] + remainder
			}
		}

		if sum > 9 {
			remainder = 1
			sum = sum % 10
		} else {
			remainder = 0
		}
		result[step] = sum
		step--
		i--
		j--
	}

	if remainder > 0 {
		result = append([]int{remainder}, result...)
	}

	return result
}

func main() {
	nums := []int{9, 9, 9, 9, 9}
	k := 1
	fmt.Println(addToArrayForm(nums, k))
}
