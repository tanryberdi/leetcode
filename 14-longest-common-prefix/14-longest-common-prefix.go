package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}

	max := 0
	maxSt := ""
	head := 0
	st := string(strs[0][head])
	for {
		found := true
		for _, word := range strs {
			index := strings.Index(word, st)
			found = found && (index == 0)
			if !found {
				break
			}
		}

		if found {
			if len(st) > max {
				max = len(st)
				maxSt = st
			}

			head++
			if head < len(strs[0]) {
				st = st + string(strs[0][head])
			} else {
				break
			}
		} else {
			break
		}
	}

	return maxSt
}

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}
