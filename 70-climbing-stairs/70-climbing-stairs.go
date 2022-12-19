package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	if n < 4 {
		return n
	}

	step := []int{1, 2}
	for i := 2; i < n; i++ {
		step = append(step, step[i-1]+step[i-2])
	}
	return step[n-1]
}

func main() {
	n := 5
	fmt.Println(climbStairs(n))
}
