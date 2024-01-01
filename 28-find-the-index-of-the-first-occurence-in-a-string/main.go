package main

// Index returns the index of the first occurrence of find in s, or -1 if find is not present in s.
func Index(haystack, needle string) int {
	// if needle is longer than haystack, return -1
	if len(needle) > len(haystack) {
		return -1
	}

	// if haystack is needle, return 0
	if haystack == needle {
		return 0
	}

	// Loop over all indexes in s.
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		// Compare s[i:len(find)] to find.
		if haystack[i:i+len(needle)] == needle {
			// If we found it, return the index.
			return i
		}
	}

	// Return -1 if we didn't find it.
	return -1
}

func main() {
	// This is the string we want to search
	s := "hello"

	// This is the string we want to find the index of
	find := "ll"

	// Find the index of the first occurrence of find in s
	i := Index(s, find)

	// Print the result
	println(i)
}
