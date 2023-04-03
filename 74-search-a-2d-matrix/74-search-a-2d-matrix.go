package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 3))
}
