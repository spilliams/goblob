package main

import "testing"

func TestBlobParsing(t *testing.T) {
	p := NewBlobParser()
	
	cases := struct{
		name string
		input string
		expect interface{}
	}{
		{
			name: "real object",
			input: `a:6:{s:5:"label";s:7:"Comment";s:8:"settings";a:1:{s:15:"text_processing";i:1;}s:8:"required";b:1;s:7:"display";a:1:{s:7:"default";a:5:{s:5:"label";s:6:"hidden";s:4:"type";s:12:"text_default";s:6:"weight";i:0;s:8:"settings";a:0:{}s:6:"module";s:4:"text";}}s:6:"widget";a:4:{s:4:"type";s:13:"text_textarea";s:8:"settings";a:1:{s:4:"rows";i:5;}s:6:"weight";i:0;s:6:"module";s:4:"text";}s:11:"description";s:0:"";}`,
			expect: map[string]interface{}{
				"label": "Comment",
				"settings": map[string]interface{}{
					"text_processing": 1,
				},
				"required": true,
				"display": map[string]interface{}{
					"default": map[string]interface{}{
						"label": "hidden",
						"type": "text_default",
						"weight": 0,
						"settings": map[string]interface{}{},
						"module": "text",
					},
				}
				"widget": map[string]interface{}{
					"type": "text_textarea",
					"settings": map[string]interface{}{
						"rows": 5,
					},
					"weight": 0,
					"module": "text",
				}
				"description": "",
			},
		},
		{
			name: "list",
			input: `a:3{i:0;s:5:12345;i:1;i:12;i:2;b:0}`,
			expect: []interface{}{
				"12345",
				12,
				false
			},
		},
	}
	
	for _, c := range cases {
		t.Run(c.name, func (t *testing.T) {
			actual := p.ParseString(c.input)
			if actual != c.expect {
				t.Errorf("Expected parsed object to match %v. Actual %v", expected, p.ParsedObject())
			}
		})
	}
}
