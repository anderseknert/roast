package encoding

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type setComprehensionCodec struct{}

func (*setComprehensionCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*setComprehensionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	sc := *((*ast.SetComprehension)(ptr))

	stream.WriteObjectStart()

	stream.WriteObjectField("term")
	stream.WriteVal(sc.Term)
	stream.WriteMore()
	stream.WriteObjectField("body")
	stream.WriteVal(sc.Body)

	stream.WriteObjectEnd()
}
