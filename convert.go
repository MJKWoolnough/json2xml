// Package json2xml converts a JSON structure to XML.
//
// json2xml wraps each type within xml tags named after the type. For example:-
//
// An object is wrapped in <object></object>
// An array is wrapped in <array></array>
//
// When a type is a member of an object, the name of the key becomes an
// attribute on the type tag, for example: -
//
// {
// 	"Location": {
// 		"Longitude": -1.8262,
// 		"Latitude": 51.1789
// 	}
// }
//
// ...becomes...
//
// <object>
//	<object name="Location">
//		<number name="Longitude">-1.8262</number>
// 		<number name="Latitude">51.1789</number>
// 	</object>
// <object>
package json2xml

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"strconv"
)

type ttype byte

const (
	typObject ttype = iota
	typArray
	typBool
	typNumber
	typString
	typNull
)

func (t ttype) String() string {
	switch t {
	case typObject:
		return "object"
	case typArray:
		return "array"
	case typBool:
		return "boolean"
	case typNumber:
		return "number"
	case typString:
		return "string"
	case typNull:
		return "null"
	default:
		return "unknown"
	}
}

// Converter represents the ongoing conversion from JSON to XML
type Converter struct {
	decoder *json.Decoder
	types   []ttype
	data    *string
}

// Tokens provides a JSON converter that implements the xml.TokenReader
// interface
func Tokens(j *json.Decoder) *Converter {
	return &Converter{
		decoder: j,
	}
}

// Token gets a xml.Token from the Converter, as per the xml.TokenReader
// interface
func (c *Converter) Token() (xml.Token, error) {
	if len(c.types) > 0 {
		switch c.types[len(c.types)-1] {
		case typObject, typArray:
		default:
			if c.data != nil {
				token := xml.CharData(*c.data)
				c.data = nil
				return token, nil
			}
			return c.outputEnd(), nil
		}
	}
	var keyName *string
	for {
		token, err := c.decoder.Token()
		if err != nil {
			return nil, err
		}
		switch token := token.(type) {
		case json.Delim:
			switch token {
			case '{':
				return c.outputStart(typObject, keyName), nil
			case '[':
				return c.outputStart(typArray, keyName), nil
			case '}', ']':
				return c.outputEnd(), nil
			}
		case bool:
			if token {
				return c.outputType(typBool, &cTrue, keyName), nil
			}
			return c.outputType(typBool, &cFalse, keyName), nil
		case float64:
			number := strconv.FormatFloat(token, 'f', -1, 64)
			return c.outputType(typNumber, &number, keyName), nil
		case json.Number:
			return c.outputType(typNumber, (*string)(&token), keyName), nil
		case string:
			if len(c.types) > 0 && c.types[len(c.types)-1] == typObject && keyName == nil {
				keyName = &token
			} else {
				return c.outputType(typString, &token, keyName), nil
			}
		case nil:
			return c.outputType(typNull, nil, keyName), nil
		}
	}
}

func (c *Converter) outputType(typ ttype, data *string, keyName *string) xml.Token {
	c.data = data
	return c.outputStart(typ, keyName)
}

func (c *Converter) outputStart(typ ttype, keyName *string) xml.Token {
	c.types = append(c.types, typ)
	var attr []xml.Attr
	if keyName != nil {
		attr = []xml.Attr{
			xml.Attr{
				Name: xml.Name{
					Local: "name",
				},
				Value: *keyName,
			},
		}
	}
	return xml.StartElement{
		Name: xml.Name{
			Local: typ.String(),
		},
		Attr: attr,
	}
}

func (c *Converter) outputEnd() xml.Token {
	typ := c.types[len(c.types)-1]
	c.types = c.types[:len(c.types)-1]
	return xml.EndElement{
		Name: xml.Name{
			Local: typ.String(),
		},
	}
}

// Convert converts JSON and sends it to the given XML encoder
func Convert(j *json.Decoder, x *xml.Encoder) error {
	c := Converter{
		decoder: j,
	}
	for {
		tk, err := c.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if err = x.EncodeToken(tk); err != nil {
			return err
		}
	}
}

var (
	cTrue  = "true"
	cFalse = "false"
)
