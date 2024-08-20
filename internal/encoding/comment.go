package encoding

import (
	"encoding/base64"
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type commentCodec struct{}

func (*commentCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*commentCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	comment := *((*ast.Comment)(ptr))

	stream.WriteObjectStart()

	if comment.Location != nil {
		stream.WriteObjectField("location")
		stream.WriteVal(comment.Location)
		stream.WriteMore()
	}

	stream.WriteObjectField("text")
	stream.WriteString(base64.StdEncoding.EncodeToString(comment.Text))

	stream.WriteObjectEnd()
}