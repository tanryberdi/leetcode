package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	return doCombinationSum(candidates, target)
}

func doCombinationSum(candidates []int, target int) [][]int {
	var res [][]int

	if len(candidates) == 0 {
		return res
	}

	t := target - candidates[0]
	if t == 0 {
		res = append(res, []int{candidates[0]})
	} else if t > 0 {
		res = doCombinationSum(candidates[1:], t)
		for i, v := range res {
			res[i] = append([]int{candidates[0]}, v...)
		}
	}

	for len(candidates) > 1 && candidates[0] == candidates[1] {
		candidates = candidates[1:]
	}

	res = append(res, doCombinationSum(candidates[1:], target)...)

	return res
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
