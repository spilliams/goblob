package parse

import (
	"fmt"
	"strconv"
	"strings"
)

type token struct {
	key    string
	length int
	value  string
}

func makeToken(s string) (token, error) {
	fmt.Printf("makeToken %s\n", s)
	split := strings.Split(s, ":")
	if len(split) < 2 {
		return token{}, fmt.Errorf("input had only %d colon args (expected at least 2)", len(split))
	}
	if len(split) == 2 {
		// second one is value
		return token{
			key:   split[0],
			value: discardLastIf(split[1], ";"),
		}, nil
	}

	// second one is length, all others are value
	t := token{
		key:   split[0],
		value: discardLastIf(strings.Join(split[2:], ":"), ";"),
	}
	length, err := strconv.Atoi(split[1])
	if err != nil {
		return t, err
	}
	t.length = length
	return t, nil
}

func (t token) String() string {
	return fmt.Sprintf("{key:%s, length:%d, value:%s}", t.key, t.length, t.value)
}

func discardLastIf(s string, t string) string {
	runes := []rune(s)
	// if t is empty we discard last regardless
	if len(t) > 0 && runes[len(runes)-1] != rune(t[0]) {
		return s
	}
	return string(runes[0 : len(runes)-1])
}
