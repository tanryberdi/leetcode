package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	pmap := make(map[string]string)
	wmap := make(map[string]string)
	for i := 0; i < len(pattern); i++ {
		p, w := string(pattern[i]), words[i]
		if pmap[p] == "" && wmap[w] == "" {
			pmap[p], wmap[w] = w, p
		} else if pmap[p] != w || wmap[w] != p {
			return false
		}
	}

	return true
}

func main() {
	pattern := "abba"
	str := "dog cat cat fish"
	fmt.Println(wordPattern(pattern, str))
}
