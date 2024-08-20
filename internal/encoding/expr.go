package encoding

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type exprCodec struct{}

func (*exprCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*exprCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	expr := *((*ast.Expr)(ptr))

	stream.WriteObjectStart()

	hasWritten := false

	if expr.Location != nil {
		stream.WriteObjectField("location")
		stream.WriteVal(expr.Location)

		hasWritten = true
	}

	if expr.Negated {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField("negated")
		stream.WriteBool(expr.Negated)

		hasWritten = true
	}

	if expr.Generated {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField("generated")
		stream.WriteBool(expr.Generated)

		hasWritten = true
	}

	if len(expr.With) > 0 {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField("with")
		stream.WriteArrayStart()

		for i, with := range expr.With {
			if i > 0 {
				stream.WriteMore()
			}

			stream.WriteVal(with)
		}

		stream.WriteArrayEnd()

		hasWritten = true
	}

	if expr.Terms != nil {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField("terms")

		switch t := expr.Terms.(type) {
		case *ast.Term:
			stream.WriteVal(t)
		case []*ast.Term:
			writeTermsArray(stream, t)
		case *ast.SomeDecl:
			stream.WriteVal(t)
		case *ast.Every:
			stream.WriteVal(t)
		}
	}

	stream.WriteObjectEnd()
}