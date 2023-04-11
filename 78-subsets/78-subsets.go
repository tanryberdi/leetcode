package main

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	var dfs func(int, []int)
	dfs = func(index int, path []int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), path...))
			return
		}
		dfs(index+1, path)
		dfs(index+1, append(path, nums[index]))
	}
	dfs(0, []int{})
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
