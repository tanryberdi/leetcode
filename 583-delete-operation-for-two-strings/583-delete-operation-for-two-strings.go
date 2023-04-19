package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range word1 {
		for j, c2 := range word2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	lcs := dp[m][n]
	return m + n - lcs*2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	word1 := "a"
	word2 := "b"
	fmt.Println(minDistance(word1, word2))
}
