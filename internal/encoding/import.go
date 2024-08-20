package encoding

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type importCodec struct{}

func (*importCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*importCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	imp := *((*ast.Import)(ptr))

	stream.WriteObjectStart()

	if imp.Location != nil {
		stream.WriteObjectField("location")
		stream.WriteVal(imp.Location)
	}

	if imp.Path != nil {
		if imp.Location != nil {
			stream.WriteMore()
		}

		stream.WriteObjectField("path")
		stream.WriteVal(imp.Path)

		if imp.Alias != "" {
			stream.WriteMore()
			stream.WriteObjectField("alias")
			stream.WriteVal(imp.Alias)
		}
	}

	stream.WriteObjectEnd()
}
