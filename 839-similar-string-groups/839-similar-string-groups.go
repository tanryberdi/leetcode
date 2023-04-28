package main

import "fmt"

type UnionFind struct {
	parent []int
	size   []int
	count  int
}

func numSimilarGroups(strs []string) int {
	uf := NewUnionFind(len(strs))
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if isSimilar(strs[i], strs[j]) {
				uf.Union(i, j)
			}
		}
	}
	return uf.Count()
}

func isSimilar(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
			if diff > 2 {
				return false
			}
		}
	}
	return true
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return
	}
	if uf.size[rootX] < uf.size[rootY] {
		rootX, rootY = rootY, rootX
	}
	uf.parent[rootY] = rootX
	uf.size[rootX] += uf.size[rootY]
	uf.count--
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func main() {
	strs := []string{"tars", "rats", "arts", "star"}
	fmt.Println(numSimilarGroups(strs))
}
