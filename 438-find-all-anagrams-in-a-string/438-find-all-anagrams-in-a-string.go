package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var res []int
	if len(s) < len(p) {
		return res
	}
	var pMap [26]int
	for i := 0; i < len(p); i++ {
		pMap[p[i]-'a']++
	}
	var sMap [26]int
	for i := 0; i < len(p); i++ {
		sMap[s[i]-'a']++
	}
	if sMap == pMap {
		res = append(res, 0)
	}
	for i := 1; i < len(s)-len(p)+1; i++ {
		sMap[s[i-1]-'a']--
		sMap[s[i+len(p)-1]-'a']++
		if sMap == pMap {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}
