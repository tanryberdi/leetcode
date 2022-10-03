package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	newArr := make([]int, len(arr))

	copy(newArr, arr)

	sort.Ints(newArr)

	ranks := make(map[int]int)

	rank := 0
	for i := 0; i < len(newArr); i++ {
		if _, ok := ranks[newArr[i]]; !ok {
			rank++
			ranks[newArr[i]] = rank
			continue
		}
		ranks[newArr[i]] = rank
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = ranks[arr[i]]
	}

	return arr
}

func main() {
	arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	fmt.Println(arrayRankTransform(arr))
}
