package json2xml

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"
)

func TestBasic(t *testing.T) {
	var buf strings.Builder
	for n, test := range []struct {
		Input, Output string
	}{
		{"{}", "<object></object>"},
		{"[]", "<array></array>"},
		{"123", "<number>123</number>"},
		{"123.456", "<number>123.456</number>"},
		{"true", "<boolean>true</boolean>"},
		{"false", "<boolean>false</boolean>"},
		{"\"A\"", "<string>A</string>"},
		{"\"Hello, World\"", "<string>Hello, World</string>"},
		{"\"Hello,\\nWorld\"", "<string>Hello,\nWorld</string>"},
		{"null", "<null></null>"},
	} {
		x := xml.NewEncoder(&buf)
		if err := Convert(json.NewDecoder(strings.NewReader(test.Input)), x); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
			continue
		}
		x.Flush()
		output := buf.String()
		buf.Reset()
		if output != test.Output {
			t.Errorf("test %d: expecting %q, got %q", n+1, test.Output, output)
		}
	}
}

func TestComplex(t *testing.T) {
	var buf strings.Builder
	for n, test := range []struct {
		Input, Output string
	}{
		{"{\"Name1\":\"String1\"}", "<object><string name=\"Name1\">String1</string></object>"},
		{"[\"Name1\",\"String1\"]", "<array><string>Name1</string><string>String1</string></array>"},
		{"[{\"A\":[{\"B\":3.14159,\"C\":null},\"D\",\"E\",null,1.234],\"F\":123},\"G\"]", "<array><object><array name=\"A\"><object><number name=\"B\">3.14159</number><null name=\"C\"></null></object><string>D</string><string>E</string><null></null><number>1.234</number></array><number name=\"F\">123</number></object><string>G</string></array>"},
	} {
		x := xml.NewEncoder(&buf)
		if err := Convert(json.NewDecoder(strings.NewReader(test.Input)), x); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
			continue
		}
		x.Flush()
		output := buf.String()
		buf.Reset()
		if output != test.Output {
			t.Errorf("test %d: expecting %q, got %q", n+1, test.Output, output)
		}
	}
}
