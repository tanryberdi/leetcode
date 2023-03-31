package main

import "fmt"

func letterCasePermutation(s string) []string {
	var res []string
	var dfs func(s string, i int)
	dfs = func(s string, i int) {
		if i == len(s) {
			res = append(res, s)
			return
		}
		dfs(s, i+1)
		if s[i] >= 'a' && s[i] <= 'z' {
			dfs(s[:i]+string(s[i]-'a'+'A')+s[i+1:], i+1)
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			dfs(s[:i]+string(s[i]-'A'+'a')+s[i+1:], i+1)
		}
	}
	dfs(s, 0)
	return res
}

func main() {
	s := "a1b2"
	fmt.Println(letterCasePermutation(s))
}
