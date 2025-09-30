# json2xml

[![CI](https://github.com/MJKWoolnough/json2xml/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/json2xml/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/json2xml.svg)](https://pkg.go.dev/vimagination.zapto.org/json2xml)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/json2xml)](https://goreportcard.com/report/vimagination.zapto.org/json2xml)

--
    import "vimagination.zapto.org/json2xml"

Package json2xml converts a JSON structure to XML.

## Highlights

 - Safely converts JSON to XML with type-based tags.
 - Object key names are stored as a name attribute.

## Usage

```go
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"vimagination.zapto.org/json2xml"
)

func main() {
	var buf strings.Builder

	jsonData := `` +
		`[
	{
		"A": [
			{
				"B": 3.14159,
				"C": null
			},
			"D",
			"E",
			null,
			1.234
		],
		"F": 123
	},
	"G"
]`

	x := xml.NewEncoder(&buf)

	x.Indent("", "\t")

	if err := json2xml.Convert(json.NewDecoder(strings.NewReader(jsonData)), x); err != nil {
		fmt.Printf("unexpected error: %s\n", err)

		return
	}

	x.Flush()

	fmt.Println(buf.String())

	// Output:
	// <array>
	//	<object>
	//		<array name="A">
	//			<object>
	//				<number name="B">3.14159</number>
	//				<null name="C"></null>
	//			</object>
	//			<string>D</string>
	//			<string>E</string>
	//			<null></null>
	//			<number>1.234</number>
	//		</array>
	//		<number name="F">123</number>
	//	</object>
	//	<string>G</string>
	// </array>
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/json2xml
