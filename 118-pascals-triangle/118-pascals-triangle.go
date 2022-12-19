package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}

	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	result := make([][]int, numRows)
	for i := range result {
		result[i] = make([]int, i+1)
	}

	result[0][0] = 1
	result[1][0] = 1
	result[1][1] = 1

	for i := 2; i < numRows; i++ {
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}

func main() {
	numRows := 4
	fmt.Println(generate(numRows))
}
