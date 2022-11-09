package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1, 1}
	}

	result := make([][]int, rowIndex+1)
	for i := range result {
		result[i] = make([]int, i+1)
	}

	result[0][0] = 1
	result[1][0] = 1
	result[1][1] = 1

	for i := 2; i <= rowIndex; i++ {
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result[rowIndex]
}

func main() {
	rowIndex := 3
	fmt.Println(getRow(rowIndex))
}
