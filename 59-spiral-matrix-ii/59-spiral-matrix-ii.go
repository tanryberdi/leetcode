package main

import "fmt"

func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	i := 0
	j := 0
	step := 0
	direction := 0
	for {
		bb := false

		if (i < n) && (i >= 0) && (j < n) && (j >= 0) && result[i][j] == 0 {
			step++
			result[i][j] = step
			bb = true
		}

		if direction == 0 && ((j+1 == n) || (result[i][j+1] > 0)) {
			direction = 1
		} else if direction == 1 && ((i+1 == n) || (result[i+1][j] > 0)) {
			direction = 2
		} else if direction == 2 && ((j-1 == -1) || (result[i][j-1] > 0)) {
			direction = 3
		} else if direction == 3 && ((i-1 == -1) || result[i-1][j] > 0) {
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

		if !bb {
			break
		}
	}

	return result
}

func main() {
	n := 1
	fmt.Println(generateMatrix(n))
}
