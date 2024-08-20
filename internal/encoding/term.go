package encoding

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type termCodec struct{}

func (*termCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*termCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	term := *((*ast.Term)(ptr))

	stream.WriteObjectStart()

	if term.Location != nil {
		stream.WriteObjectField("location")
		stream.WriteVal(term.Location)
	}

	if term.Value != nil {
		if term.Location != nil {
			stream.WriteMore()
		}

		stream.WriteObjectField("type")
		stream.WriteString(ast.TypeName(term.Value))
		stream.WriteMore()
		stream.WriteObjectField("value")
		stream.WriteVal(term.Value)
	}

	stream.WriteObjectEnd()
}
