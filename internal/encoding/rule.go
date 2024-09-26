package encoding

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"

	"github.com/open-policy-agent/opa/ast"
)

type ruleCodec struct{}

func (*ruleCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*ruleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	rule := *((*ast.Rule)(ptr))

	stream.WriteObjectStart()

	hasWritten := false

	if rule.Location != nil {
		stream.WriteObjectField(strLocation)
		stream.WriteVal(rule.Location)

		hasWritten = true
	}

	if len(rule.Annotations) > 0 {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField(strAnnotations)
		stream.WriteArrayStart()

		for i, ann := range rule.Annotations {
			if i > 0 {
				stream.WriteMore()
			}

			stream.WriteVal(ann)
		}

		stream.WriteArrayEnd()

		hasWritten = true
	}

	if rule.Default {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField(strDefault)
		stream.WriteBool(rule.Default)

		hasWritten = true
	}

	if rule.Head != nil {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField(strHead)
		stream.WriteObjectStart()

		hasWrittenHead := false

		if rule.Head.Location != nil {
			stream.WriteObjectField(strLocation)
			stream.WriteVal(rule.Head.Location)

			hasWrittenHead = true
		}

		if rule.Head.Reference != nil {
			if hasWrittenHead {
				stream.WriteMore()
			}

			stream.WriteObjectField(strRef)
			stream.WriteVal(rule.Head.Reference)

			hasWrittenHead = true
		}

		if len(rule.Head.Args) > 0 {
			if hasWrittenHead {
				stream.WriteMore()
			}

			stream.WriteObjectField(strArgs)
			writeTermsArray(stream, rule.Head.Args)

			hasWrittenHead = true
		}

		if rule.Head.Assign {
			if hasWrittenHead {
				stream.WriteMore()
			}

			stream.WriteObjectField(strAssign)
			stream.WriteBool(rule.Head.Assign)

			hasWrittenHead = true
		}

		if rule.Head.Key != nil {
			if hasWrittenHead {
				stream.WriteMore()
			}

			stream.WriteObjectField(strKey)
			stream.WriteVal(rule.Head.Key)

			hasWrittenHead = true
		}

		if rule.Head.Value != nil {
			if hasWrittenHead {
				stream.WriteMore()
			}

			stream.WriteObjectField(strValue)
			stream.WriteVal(rule.Head.Value)
		}

		stream.WriteObjectEnd()

		hasWritten = true
	}

	if !isBodyGenerated(&rule) {
		if hasWritten {
			stream.WriteMore()
		}

		stream.WriteObjectField(strBody)
		stream.WriteVal(rule.Body)
	}

	if rule.Else != nil {
		stream.WriteMore()
		stream.WriteObjectField(strElse)
		stream.WriteVal(rule.Else)
	}

	stream.WriteObjectEnd()
}

func isBodyGenerated(rule *ast.Rule) bool {
	if rule.Default {
		return true
	}

	if len(rule.Body) == 0 {
		return true
	}

	if rule.Head == nil {
		return false
	}

	if rule.Body[0] != nil && rule.Body[0].Location == rule.Location {
		return true
	}

	if rule.Body[0] != nil && rule.Head.Value != nil && rule.Body[0].Location == rule.Head.Value.Location {
		return true
	}

	if rule.Head.Key != nil &&
		rule.Body[0].Location.Row == rule.Head.Key.Location.Row &&
		rule.Body[0].Location.Col < rule.Head.Key.Location.Col {
		// This is a quirk in the original AST — the generated body will have a location
		// set before the key, i.e. "message"
		return true
	}

	return false
}
