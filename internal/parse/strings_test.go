package parse

import "testing"

func TestMatchBrace(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		start  int
		open   rune
		close  rune
		expect int
	}{
		{
			name:   "first",
			input:  "a:3:{asdbasdb}",
			start:  4,
			open:   '{',
			close:  '}',
			expect: 13,
		},
		{
			name:   "nested",
			input:  "{aklsjdhf{aksldhjf}{}}alskdjfh",
			start:  0,
			open:   '{',
			close:  '}',
			expect: 21,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := matchBrace(c.input, c.start, c.open, c.close)
			if actual != c.expect {
				t.Errorf("brace mismatch! Expected %v at %d, actually at %d (%v is at %d)", c.close, c.expect, actual, c.input[actual], actual)
			}
		})
	}
}

func TestReadUntil(t *testing.T) {
	cases := []struct {
		name         string
		input        string
		start        int
		until        rune
		expectIndex  int
		expectString string
	}{
		{
			name:         "simple",
			input:        "asdf:fddsa:lkjh",
			start:        5,
			until:        ':',
			expectIndex:  10,
			expectString: "fddsa:",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualIndex, actualString := readUntil(c.input, c.start, c.until)
			if actualIndex != c.expectIndex {
				t.Errorf("index mismatch: expected %d, actual %d", c.expectIndex, actualIndex)
			}
			if actualString != c.expectString {
				t.Errorf("substring mismatch: expected %s, actual %s", c.expectString, actualString)
			}
		})
	}
}
