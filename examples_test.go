package json2xml_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"vimagination.zapto.org/json2xml"
)

func Example() {
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
