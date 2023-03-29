package main

import "fmt"

func orangesRotting(grid [][]int) int {
	var (
		rottenOranges int
		rotten        bool
	)
	for {
		rotten = false
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 2 {
					if i-1 >= 0 && grid[i-1][j] == 1 {
						grid[i-1][j] = 3
						rotten = true
					}
					if i+1 < len(grid) && grid[i+1][j] == 1 {
						grid[i+1][j] = 3
						rotten = true
					}
					if j-1 >= 0 && grid[i][j-1] == 1 {
						grid[i][j-1] = 3
						rotten = true
					}
					if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
						grid[i][j+1] = 3
						rotten = true
					}
				}
			}
		}
		if !rotten {
			break
		}
		rottenOranges++
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 3 {
					grid[i][j] = 2
				}
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return rottenOranges
}

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(orangesRotting(grid))
}
