package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return []int{1}
	}

	result := make([]int, 0)

	i := 0
	j := 0
	direction := 0

	for {
		bb := false

		if (i < len(matrix)) && (i >= 0) && (j < len(matrix[0])) && (j >= 0) && matrix[i][j] != 200 {
			result = append(result, matrix[i][j])
			matrix[i][j] = 200
			bb = true
		}

		if direction == 0 && ((j+1 == len(matrix[0])) || (matrix[i][j+1] == 200)) {
			direction = 1
		} else if direction == 1 && ((i+1 == len(matrix)) || (matrix[i+1][j] == 200)) {
			direction = 2
		} else if direction == 2 && ((j-1 == -1) || (matrix[i][j-1] == 200)) {
			direction = 3
		} else if direction == 3 && ((i-1 == -1) || matrix[i-1][j] == 200) {
			direction = 0
		}

		switch direction {
		case 0:
			j++
		case 1:
			i++
		case 2:
			j--
		default:
			i--
		}

		if i == len(matrix) || j == len(matrix[0]) || i < 0 || j < 0 {
			bb = false
		}

		if !bb {
			break
		}

	}

	return result
}

func main() {
	matrix := [][]int{{3}, {2}}
	fmt.Println(spiralOrder(matrix))
}
