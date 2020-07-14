package parse

import (
	"reflect"
	"testing"
)

func TestReadList(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect []token
	}{
		{
			name:  "integers",
			input: "i:0;i:1;i:2;",
			expect: []token{
				{key: "i", length: 0, value: "0"},
				{key: "i", length: 0, value: "1"},
				{key: "i", length: 0, value: "2"},
			},
		},
		{
			name:  "mix",
			input: "i:0;s:5:\"asdfg\";b:1;",
			expect: []token{
				{key: "i", length: 0, value: "0"},
				{key: "s", length: 5, value: "\"asdfg\""},
				{key: "b", length: 0, value: "1"},
			},
		},
		{
			name:  "array",
			input: "a:1:{i:0;s:2:\"ab\";}a:1:{s:1:\"a\";s:0:\"\";}",
			expect: []token{
				{key: "a", length: 1, value: "{i:0;s:2:\"ab\";}"},
				{key: "a", length: 1, value: "{s:1:\"a\";s:0:\"\";}"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := readList(c.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("token mismatch. Expected: %v Actual: %v", c.expect, actual)
			}
		})
	}
}

func TestBlobParsing(t *testing.T) {
	p := NewBlobParser()

	cases := []struct {
		name   string
		input  string
		expect interface{}
	}{
		{
			name:   "integer",
			input:  `i:0;`,
			expect: 0,
		},
		{
			name:   "string",
			input:  `s:4:"true";`,
			expect: "true",
		},
		{
			name:   "empty_string",
			input:  `s:0:"";`,
			expect: "",
		},
		{
			name:   "false",
			input:  `b:0;`,
			expect: false,
		},
		{
			name:   "true",
			input:  `b:1;`,
			expect: true,
		},
		{
			name:  "simple_list",
			input: `a:3:{i:0;s:5:"12345";i:1;i:12;i:2;b:0;}`,
			expect: []interface{}{
				"12345",
				12,
				false,
			},
		},
		{
			name:  "nested_list",
			input: `a:2:{i:0;a:0:{}i:1;a:2:{i:0;b:0;i:1;s:2:"cd";i:2;a:1:{i:0;s:3:"abc";}}}`,
			expect: []interface{}{
				map[string]interface{}{},
				[]interface{}{
					false,
					"cd",
					[]interface{}{
						"abc",
					},
				},
			},
		},
		{
			name:  "nested_map",
			input: `a:5:{s:3:"int";i:1;s:4:"bool";b:0;s:3:"str";s:3:"foo";s:4:"list";a:4:{i:0;i:0;i:1;i:1;i:2;i:2;i:3;b:0;}s:3:"map";a:1:{s:3:"int";i:0;}}`,
			expect: map[string]interface{}{
				"int":  1,
				"bool": false,
				"str":  "foo",
				"list": []interface{}{
					0,
					1,
					2,
					false,
				},
				"map": map[string]interface{}{
					"int": 0,
				},
			},
		},
		{
			name:  "real object",
			input: `a:6:{s:5:"label";s:7:"Comment";s:8:"settings";a:1:{s:15:"text_processing";i:1;}s:8:"required";b:1;s:7:"display";a:1:{s:7:"default";a:5:{s:5:"label";s:6:"hidden";s:4:"type";s:12:"text_default";s:6:"weight";i:0;s:8:"settings";a:0:{}s:6:"module";s:4:"text";}}s:6:"widget";a:4:{s:4:"type";s:13:"text_textarea";s:8:"settings";a:1:{s:4:"rows";i:5;}s:6:"weight";i:0;s:6:"module";s:4:"text";}s:11:"description";s:0:"";}`,
			expect: map[string]interface{}{
				"label": "Comment",
				"settings": map[string]interface{}{
					"text_processing": 1,
				},
				"required": true,
				"display": map[string]interface{}{
					"default": map[string]interface{}{
						"label":    "hidden",
						"type":     "text_default",
						"weight":   0,
						"settings": map[string]interface{}{},
						"module":   "text",
					},
				},
				"widget": map[string]interface{}{
					"type": "text_textarea",
					"settings": map[string]interface{}{
						"rows": 5,
					},
					"weight": 0,
					"module": "text",
				},
				"description": "",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			parseErr := p.Parse(c.input)
			if parseErr != nil {
				t.Error(parseErr)
			}
			actual := p.ParsedObject()
			// need reflect on account of all the interface{}s
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("Expected parsed object to match %v (%T). Actual %v (%T)", c.expect, c.expect, actual, actual)
			}
		})
	}
}
