package main

import "fmt"

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			if dp[j]*dp[i-j] > max {
				max = dp[j] * dp[i-j]
			}
		}
		dp[i] = max
	}
	return dp[n]
}

func main() {
	n := 10
	fmt.Println("Output: ", integerBreak(n))
}
