package main

import "fmt"

func combine(n int, k int) [][]int {
	if k == 0 {
		return [][]int{{}}
	}
	if n == 0 {
		return [][]int{}
	}
	var res [][]int
	for i := n; i >= 1; i-- {
		for _, v := range combine(i-1, k-1) {
			res = append(res, append(v, i))
		}
	}
	return res
}

func main() {
	fmt.Println(combine(4, 2))
}
