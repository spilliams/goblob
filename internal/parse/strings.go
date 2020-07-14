package parse

// matchBrace reads a given string and returns the index of the closing rune
// matching the opening rune at the starting index.
// Returns -1 if it doesn't find a closing rune.
func matchBrace(s string, start int, open rune, close rune) int {
	openCount := 0
	for i := start; i < len(s); i++ {
		r := rune(s[i])
		if r == open {
			openCount++
		}
		if r == close {
			openCount--
		}
		if openCount == 0 {
			return i
		}
	}
	return -1
}

// readUntil reads a given string starting at an index until it reaches a
// certain rune (or the end of the string). Returns the number of characters
// read as well as the substring between start and end (inclusive)
func readUntil(s string, start int, until rune) (int, string) {
	build := ""
	for i := start; i < len(s); i++ {
		build += string(s[i])
		if rune(s[i]) == until {
			return i, build
		}
	}
	return len(s) - start, build
}
