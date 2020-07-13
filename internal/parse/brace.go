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
