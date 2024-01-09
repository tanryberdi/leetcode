package main

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var m = make(map[rune]int)
	for _, c := range s {
		m[c]++
	}

	for _, c := range t {
		m[c]--
		if m[c] < 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "rat"
	t := "car"
	println(isAnagram(s, t))
}
