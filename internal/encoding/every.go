package encoding

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type everyCodec struct{}

func (*everyCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*everyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	every := *((*ast.Every)(ptr))

	stream.WriteObjectStart()

	if every.Location != nil {
		stream.WriteObjectField("location")
		stream.WriteVal(every.Location)
		stream.WriteMore()
	}

	stream.WriteObjectField("key")
	stream.WriteVal(every.Key)
	stream.WriteMore()

	stream.WriteObjectField("value")
	stream.WriteVal(every.Value)
	stream.WriteMore()

	stream.WriteObjectField("domain")
	stream.WriteVal(every.Domain)
	stream.WriteMore()

	stream.WriteObjectField("body")
	stream.WriteVal(every.Body)

	stream.WriteObjectEnd()
}
