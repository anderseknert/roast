package encoding

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type packageCodec struct{}

func (*packageCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*packageCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	pkg := *((*ast.Package)(ptr))

	stream.WriteObjectStart()

	if pkg.Location != nil {
		stream.WriteObjectField("location")
		stream.WriteVal(pkg.Location)
	}

	if pkg.Path != nil {
		if pkg.Location != nil {
			stream.WriteMore()
		}

		stream.WriteObjectField("path")
		stream.WriteVal(pkg.Path)
	}

	stream.WriteObjectEnd()
}
