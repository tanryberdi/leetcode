package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var tmp []int
	dfs(candidates, target, 0, &tmp, &result)
	return result
}

func dfs(candidates []int, target, start int, tmp *[]int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*result = append(*result, append([]int{}, *tmp...))
		return
	}
	for i := start; i < len(candidates); i++ {
		*tmp = append(*tmp, candidates[i])
		dfs(candidates, target-candidates[i], i, tmp, result)
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
