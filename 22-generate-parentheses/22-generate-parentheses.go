package main

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	generate("", n, n, &res)
	return res
}

func generate(s string, left, right int, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}
	if left > 0 {
		generate(s+"(", left-1, right, res)
	}
	if right > left {
		generate(s+")", left, right-1, res)
	}
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}
