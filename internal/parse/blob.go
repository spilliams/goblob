package parse

import (
	"fmt"
	"strconv"
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
	bp.parsedObj, err = bp.parse(s)
	return err
}

// ParsedObject returns the receiver's most recently parsed object
func (bp *BlobParser) ParsedObject() interface{} {
	return bp.parsedObj
}

func (bp *BlobParser) parse(s string) (interface{}, error) {
	t, err := makeToken(s)
	if err != nil {
		return nil, err
	}
	return parseToken(t)
}

func parseToken(t token) (interface{}, error) {
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
		return nil, fmt.Errorf("unknown data type '%s' in input", t.key)
	}
}

const (
	tokenArrayKey  = "a"
	tokenBoolKey   = "b"
	tokenIntKey    = "i"
	tokenStringKey = "s"
)

func readList(s string) ([]token, error) {
	tokenList := []token{}
	for cursor := 0; cursor < len(s); cursor++ {
		switch s[cursor] {
		case 'a':
			open, _ := readUntil(s, cursor, '{')
			close := matchBrace(s, open, '{', '}')
			t, err := makeToken(s[cursor : close+1])
			if err != nil {
				return nil, err
			}
			tokenList = append(tokenList, t)
			cursor = close
			continue
		default:
			end, value := readUntil(s, cursor, ';')
			t, err := makeToken(value)
			if err != nil {
				return nil, err
			}
			tokenList = append(tokenList, t)
			cursor = end
			if end == -1 {
				return tokenList, nil
			}
		}
	}
	return tokenList, nil
}

// ex: a:N:{i:0;...}
func parseList(listToken token) ([]interface{}, error) {
	// remove the curlies
	readValue := listToken.value[1 : len(listToken.value)-1]
	tokenChildren, err := readList(readValue)
	if err != nil {
		return nil, err
	}
	listValues := []interface{}{}

	for i, t := range tokenChildren {
		if i%2 == 0 {
			continue
		}
		childValue, err := parseToken(t)
		if err != nil {
			return nil, err
		}
		listValues = append(listValues, childValue)
	}
	return listValues, nil
}

// ex: a:N:{s:3:"key";...}
func parseMap(mapToken token) (map[string]interface{}, error) {
	// remove the curlies
	readValue := mapToken.value[1 : len(mapToken.value)-1]
	tokenChildren, err := readList(readValue)
	if err != nil {
		return nil, err
	}
	mapValues := map[string]interface{}{}

	for i := 0; i < len(tokenChildren); i += 2 {
		keyToken := tokenChildren[i]
		valueToken := tokenChildren[i+1]

		childValue, err := parseToken(valueToken)
		if err != nil {
			return nil, err
		}
		// get rid of the quotes
		key := keyToken.value[1 : len(keyToken.value)-1]
		if keyToken.length != len(key) {
			return nil, fmt.Errorf("key string is the wrong length (expected %v, actual %v)", keyToken.length, len(key))
		}
		mapValues[key] = childValue
	}
	return mapValues, nil
}

// ex: s:4:"abcd"
func parseString(t token) (string, error) {
	// get rid of the quotes
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
