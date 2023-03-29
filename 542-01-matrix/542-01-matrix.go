package main

import "fmt"

func updateMatrix(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				continue
			}
			matrix[i][j] = minDistance(matrix, i, j)
		}
	}
	return matrix
}

func minDistance(matrix [][]int, i, j int) int {
	min := 100000
	for k := 0; k < len(matrix); k++ {
		for l := 0; l < len(matrix[k]); l++ {
			if matrix[k][l] == 0 {
				distance := abs(i-k) + abs(j-l)
				if distance < min {
					min = distance
				}
			}
		}
	}
	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	fmt.Println(updateMatrix(mat))
}
