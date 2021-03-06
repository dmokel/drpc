package codec

import "io"

// Header ...
type Header struct {
	ServiceMethod string
	Seq           uint64
	Error         string
}

// Codec ...
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

// NewCodecFunc ...
type NewCodecFunc func(io.ReadWriteCloser) Codec

// Type ...
type Type string

const (
	// GobType ...
	GobType Type = "application/gob"
	// JSONType ...
	JSONType Type = "application/json"
)

// NewCodecFuncMap ...
var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
	NewCodecFuncMap[JSONType] = NewJSONCodec
}
