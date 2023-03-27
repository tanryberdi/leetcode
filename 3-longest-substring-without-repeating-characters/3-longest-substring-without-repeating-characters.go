package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	max := 0
	for i, j := 0, 0; j < len(s); j++ {
		if _, ok := m[s[j]]; ok {
			i = maxInt(i, m[s[j]]+1)
		}
		m[s[j]] = j
		max = maxInt(max, j-i+1)
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "abcabcbb"
	println(lengthOfLongestSubstring(s))
}
