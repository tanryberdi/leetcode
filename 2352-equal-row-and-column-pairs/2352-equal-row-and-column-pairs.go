package main

import "fmt"

func equalPairs(grid [][]int) int {
	numOfPairs := 0

	// trace for rows
	for i := 0; i < len(grid); i++ {
		// we are checking whether row i is same column j
		for j := 0; j < len(grid); j++ {
			equal := true
			for k := 0; k < len(grid); k++ {
				equal = equal && (grid[i][k] == grid[k][j])
				if !equal {
					break
				}
			}
			if equal {
				numOfPairs++
			}
		}
	}

	return numOfPairs
}

func main() {
	// grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
	grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	fmt.Println(equalPairs(grid))
}
