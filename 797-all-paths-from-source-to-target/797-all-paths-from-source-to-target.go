package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	var path []int
	dfs(graph, 0, &path, &res)
	return res
}

func dfs(graph [][]int, node int, path *[]int, res *[][]int) {
	*path = append(*path, node)
	if node == len(graph)-1 {
		*res = append(*res, append([]int{}, *path...))
	} else {
		for _, next := range graph[node] {
			dfs(graph, next, path, res)
		}
	}
	*path = (*path)[:len(*path)-1]
}

func main() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}
