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

// ParseString parses a string-encoded blob
// Empty arrays will parse as an empty map[string]interface{}
func (bp *BlobParser) Parse(s string) error {
	split := strings.Split(s, ":")
	switch split[0] {
	case "a":
		if len(split) < 3 {
			return fmt.Errorf("input indicated array but it had %d colon-separated elements (expected >=3)", len(split))
		}
		if len(split[2]) < 2 {
			return fmt.Errorf("input indicated array but third colon arg was unexpectedly short")
		}
		runes := []rune(split[2])
		arrayType := string(runes[1])
		switch arrayType {
		case "}":
			// empty array input
			bp.parsedObj = map[string]interface{}{}
			return nil
		case "i":
			return bp.parseList(s)
		case "s":
			return bp.parseMap(s)
		default:
			return fmt.Errorf("input indicated array but array type is unrecognized (%c)", split[2][1])
		}
	case "s":
		return bp.parseString(s)
	case "i":
		return bp.parseInteger(s)
	case "b":
		return bp.parseBoolean(s)
	default:
		return fmt.Errorf("unknown data type '%c' in input", s[0])
	}
}

func (bp *BlobParser) parseList(s string) error {
	// TODO implement me
	return fmt.Errorf("not implemented list")
}

func (bp *BlobParser) parseMap(s string) error {
	// TODO implement me
	return fmt.Errorf("not implemented map")
}

// ex: s:4:"abcd";
func (bp *BlobParser) parseString(s string) error {
	split := strings.Split(s, ":")
	if len(split) != 3 {
		return fmt.Errorf("input indicated string but it had the wrong number of colon args (expected 3, actual %d)", len(split))
	}
	runes := []rune(split[2])
	bp.parsedObj = string(runes[1 : len(runes)-2])
	return nil
}

func (bp *BlobParser) parseInteger(s string) error {
	split := strings.Split(s, ":")
	if len(split) != 2 {
		return fmt.Errorf("input indicated integer but it had the wrong number of colon args (expected 2, actual %d)", len(split))
	}
	var err error
	runes := []rune(split[1])
	bp.parsedObj, err = strconv.Atoi(string(runes[0 : len(runes)-1]))
	if err != nil {
		return err
	}
	return nil
}

func (bp *BlobParser) parseBoolean(s string) error {
	split := strings.Split(s, ":")
	if len(split) != 2 {
		return fmt.Errorf("input indicated boolean but it had the wrong number of colon args (expected 2, actuald %d)", len(split))
	}
	runes := []rune(split[1])
	value := string(runes[0])
	switch value {
	case "0":
		bp.parsedObj = false
		return nil
	case "1":
		bp.parsedObj = true
		return nil
	}
	return fmt.Errorf("unrecognized boolean value (expect 0 or 1, actual %v)", split[1])
}

// ParsedObject returns the receiver's most recently parsed object
func (bp *BlobParser) ParsedObject() interface{} {
	return bp.parsedObj
}
