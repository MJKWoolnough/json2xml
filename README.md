# json2xml
--
    import "github.com/MJKWoolnough/json2xml"

Package json2xml converts a JSON structure to XML.

json2xml wraps each type within xml tags named after the type. For example:-

An object is wrapped in <object></object> An array is wrapped in <array></array>

When a type is a member of an object, the name of the key becomes an attribute
on the type tag, for example: -

{

    "Location": {
    	"Longitude": -1.8262,
    	"Latitude": 51.1789
    }

}

...becomes...

<object>

    <object name="Location">
    	<number name="Longitude">-1.8262</number>
    	<number name="Latitude">51.1789</number>
    </object>

<object>

## Usage

```go
const (
	ErrInvalidKey   errors.Error = "invalid key type"
	ErrUnknownToken errors.Error = "unknown token type"
	ErrInvalidToken errors.Error = "invalid token"
)
```
Errors

#### func  Convert

```go
func Convert(j JSONDecoder, x XMLEncoder) error
```
Convert converts JSON and sends it to the given XML encoder

#### type Converter

```go
type Converter struct {
}
```

Converter represents the ongoing conversion from JSON to XML

#### func  Tokens

```go
func Tokens(j JSONDecoder) *Converter
```
Tokens provides a JSON converter that implements the xml.TokenReader interface

#### func (*Converter) Token

```go
func (c *Converter) Token() (xml.Token, error)
```
Token gets a xml.Token from the Converter, as per the xml.TokenReader interface

#### type JSONDecoder

```go
type JSONDecoder interface {
	Token() (json.Token, error)
}
```

JSONDecoder represents a type that gives out JSON tokens, usually implemented by
*json.Decoder

#### type XMLEncoder

```go
type XMLEncoder interface {
	EncodeToken(xml.Token) error
}
```

XMLEncoder represents a type that takes XML tokens, usually implemented by
*xml.Encoder
