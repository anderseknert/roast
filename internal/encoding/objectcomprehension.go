package encoding

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type objectComprehensionCodec struct{}

func (*objectComprehensionCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*objectComprehensionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	oc := *((*ast.ObjectComprehension)(ptr))

	stream.WriteObjectStart()

	stream.WriteObjectField("key")
	stream.WriteVal(oc.Key)
	stream.WriteMore()
	stream.WriteObjectField("value")
	stream.WriteVal(oc.Value)
	stream.WriteMore()
	stream.WriteObjectField("body")
	stream.WriteVal(oc.Body)

	stream.WriteObjectEnd()
}
