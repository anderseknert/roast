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
		stream.WriteObjectField(strLocation)
		stream.WriteVal(pkg.Location)
	}

	if pkg.Path != nil {
		if pkg.Location != nil {
			stream.WriteMore()
		}

		stream.WriteObjectField(strPath)
		stream.WriteVal(pkg.Path)
	}

	if stream.Attachment != nil {
		stream.WriteMore()
		stream.WriteObjectField(strAnnotations)
		stream.WriteVal(stream.Attachment)
	}

	stream.WriteObjectEnd()
}
