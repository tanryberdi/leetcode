package main

import "fmt"

func minimum(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	triangle[1][0] = triangle[1][0] + triangle[0][0]
	triangle[1][1] = triangle[1][1] + triangle[0][0]
	if len(triangle) == 2 {
		return minimum(triangle[1][0], triangle[1][1])
	}

	for i := 2; i < len(triangle); i++ {
		triangle[i][0] = triangle[i][0] + triangle[i-1][0]
		triangle[i][i] = triangle[i][i] + triangle[i-1][i-1]

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i][j] + minimum(triangle[i-1][j-1], triangle[i-1][j])
		}
	}

	min := triangle[len(triangle)-1][0]
	for i := 1; i < len(triangle); i++ {
		if min > triangle[len(triangle)-1][i] {
			min = triangle[len(triangle)-1][i]
		}
	}

	return min
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
