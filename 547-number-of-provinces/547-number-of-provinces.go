package main

import "fmt"

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(isConnected, visited, i)
			count++
		}
	}
	return count
}

func dfs(isConnected [][]int, visited []bool, i int) {
	for j := 0; j < len(isConnected); j++ {
		if isConnected[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(isConnected, visited, j)
		}
	}
}

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(isConnected))
}
