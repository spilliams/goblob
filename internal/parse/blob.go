package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// BlobParser represents an object that can parse blobs
type BlobParser struct {
	parsedObj interface{}
}

// NewBlobParser returns a new parser
func NewBlobParser() *BlobParser {
	return &BlobParser{}
}

// Parse parses a string-encoded blob
// Empty arrays will parse as an empty map[string]interface{}
func (bp *BlobParser) Parse(s string) error {
	var err error
	bp.parsedObj, err = parse(s)
	return err
}

// ParsedObject returns the receiver's most recently parsed object
func (bp *BlobParser) ParsedObject() interface{} {
	return bp.parsedObj
}

func parse(s string) (interface{}, error) {
	t, err := makeToken(s)
	if err != nil {
		return nil, err
	}
	switch t.key {
	case "a":
		arrayType := string(t.value[1])
		switch arrayType {
		case "}":
			// empty array input
			return map[string]interface{}{}, nil
		case "i":
			return parseList(t)
		case "s":
			return parseMap(t)
		default:
			return nil, fmt.Errorf("input indicated array but array type is unrecognized (%v)", arrayType)
		}
	case "s":
		return parseString(t)
	case "i":
		return parseInteger(t)
	case "b":
		return parseBoolean(t)
	default:
		return nil, fmt.Errorf("unknown data type '%c' in input", s[0])
	}
}

const (
	tokenArrayKey  = "a"
	tokenBoolKey   = "b"
	tokenIntKey    = "i"
	tokenStringKey = "s"
)

type token struct {
	key    string
	length int
	value  string
}

func makeToken(s string) (token, error) {
	split := strings.Split(s, ":")
	if len(split) < 2 || len(split) > 3 {
		return token{}, fmt.Errorf("input had only %d colon args (expected 2 or 3)", len(split))
	}
	runes := []rune(split[len(split)-1])
	value := string(runes[0 : len(runes)-2])
	t := token{
		key:   split[0],
		value: value,
	}
	if len(split) == 3 {
		length, err := strconv.Atoi(split[1])
		if err != nil {
			return t, err
		}
		t.length = length
	}

	return t, nil
}

// ex: a:N:{i:0;...}
func parseList(t token) ([]interface{}, error) {
	value := string(t.value[1 : len(t.value)-1])
	split := strings.Split(value, ";")
	list := []interface{}{}
	for i, v := range split {
		if i%2 == 0 {
			continue
		}
		parsedVal, err := parse(v)
		if err != nil {
			return nil, err
		}
		list = append(list, parsedVal)
	}
	return list, nil
}

// ex: a:N:{s:3:"key";...}
func parseMap(t token) (map[string]interface{}, error) {
	value := string(t.value[1 : len(t.value)-1])
	split := strings.Split(value, ";")
	dict := map[string]interface{}{}
	for i := 0; i < len(split); i += 2 {
		keyToken, err := makeToken(split[i])
		if err != nil {
			return nil, err
		}
		if keyToken.key != tokenStringKey {
			return nil, fmt.Errorf("map key was type %v (expected %v)", keyToken.key, tokenStringKey)
		}
		val, err := parse(split[i+1])
		if err != nil {
			return nil, err
		}
		dict[string(keyToken.value)] = val
	}
	return dict, nil
}

// ex: s:4:"abcd"
func parseString(t token) (string, error) {
	s := t.value[1 : len(t.value)-1]
	if t.length != len(s) {
		return "", fmt.Errorf("string is the wrong length (expected %v, actual %v)", t.length, len(s))
	}
	return s, nil
}

func parseInteger(t token) (int, error) {
	return strconv.Atoi(t.value)
}

func parseBoolean(t token) (bool, error) {
	return strconv.ParseBool(t.value)
}
