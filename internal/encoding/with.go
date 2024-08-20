package encoding

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type withCodec struct{}

func (*withCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*withCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	with := *((*ast.With)(ptr))

	stream.WriteObjectStart()

	if with.Location != nil {
		stream.WriteObjectField("location")
		stream.WriteVal(with.Location)
		stream.WriteMore()
	}

	stream.WriteObjectField("target")
	stream.WriteVal(with.Target)
	stream.WriteMore()
	stream.WriteObjectField("value")
	stream.WriteVal(with.Value)

	stream.WriteObjectEnd()
}
