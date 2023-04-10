package main

import "fmt"

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	n := len(grid)
	if n == 1 {
		return 1
	}
	grid[0][0] = 1
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0})
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, dir := range [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newX < n && newY >= 0 && newY < n && grid[newX][newY] == 0 {
				grid[newX][newY] = grid[x][y] + 1
				if newX == n-1 && newY == n-1 {
					return grid[newX][newY]
				}
				queue = append(queue, []int{newX, newY})
			}
		}
	}
	return -1
}

func main() {
	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	fmt.Println(shortestPathBinaryMatrix(grid))
}
