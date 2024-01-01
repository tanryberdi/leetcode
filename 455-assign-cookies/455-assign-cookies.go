package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi, si := 0, 0
	for gi < len(g) && si < len(s) {
		if g[gi] <= s[si] {
			gi++
		}
		si++
	}
	return gi
}

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	fmt.Print(findContentChildren(g, s))
}
