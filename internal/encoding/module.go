package encoding

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type moduleCodec struct{}

func (*moduleCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*moduleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	mod := *((*ast.Module)(ptr))

	stream.WriteObjectStart()

	hasWritten := false

	if mod.Package != nil {
		stream.WriteObjectField("package")
		stream.WriteVal(mod.Package)

		hasWritten = true
	}

	if len(mod.Imports) > 0 {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField("imports")
		stream.WriteArrayStart()

		for i, imp := range mod.Imports {
			if i > 0 {
				stream.WriteMore()
			}

			stream.WriteVal(imp)
		}

		stream.WriteArrayEnd()

		hasWritten = true
	}

	if len(mod.Annotations) > 0 {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField("annotations")
		stream.WriteArrayStart()

		for i, ann := range mod.Annotations {
			if i > 0 {
				stream.WriteMore()
			}

			stream.WriteVal(ann)
		}

		stream.WriteArrayEnd()

		hasWritten = true
	}

	if len(mod.Rules) > 0 {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField("rules")
		stream.WriteArrayStart()

		for i, rule := range mod.Rules {
			if i > 0 {
				stream.WriteMore()
			}

			stream.WriteVal(rule)
		}

		stream.WriteArrayEnd()

		hasWritten = true
	}

	if len(mod.Comments) > 0 {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField("comments")
		stream.WriteArrayStart()

		for i, comment := range mod.Comments {
			if i > 0 {
				stream.WriteMore()
			}

			stream.WriteVal(comment)
		}

		stream.WriteArrayEnd()
	}

	stream.WriteObjectEnd()
}