package main

import "fmt"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		if firstList[i][0] > secondList[j][1] {
			j++
		} else if firstList[i][1] < secondList[j][0] {
			i++
		} else {
			res = append(res, []int{max(firstList[i][0], secondList[j][0]), min(firstList[i][1], secondList[j][1])})
			if firstList[i][1] < secondList[j][1] {
				i++
			} else {
				j++
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))

}
