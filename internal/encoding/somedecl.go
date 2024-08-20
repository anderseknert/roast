package encoding

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type someDeclCodec struct{}

func (*someDeclCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*someDeclCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	some := *((*ast.SomeDecl)(ptr))

	stream.WriteObjectStart()

	if some.Location != nil {
		stream.WriteObjectField("location")
		stream.WriteVal(some.Location)
		stream.WriteMore()
	}

	stream.WriteObjectField("symbols")

	writeTermsArray(stream, some.Symbols)

	stream.WriteObjectEnd()
}
