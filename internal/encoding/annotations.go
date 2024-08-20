package encoding

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type annotationsCodec struct{}

func (*annotationsCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*annotationsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	a := *((*ast.Annotations)(ptr))

	stream.WriteObjectStart()

	if a.Location != nil {
		stream.WriteObjectField("location")
		stream.WriteVal(a.Location)
		stream.WriteMore()
	}

	stream.WriteObjectField("scope")
	stream.WriteString(a.Scope)

	if a.Title != "" {
		stream.WriteMore()
		stream.WriteObjectField("title")
		stream.WriteString(a.Title)
	}

	if a.Description != "" {
		stream.WriteMore()
		stream.WriteObjectField("description")
		stream.WriteString(a.Description)
	}

	if a.Entrypoint {
		stream.WriteMore()
		stream.WriteObjectField("entrypoint")
		stream.WriteBool(a.Entrypoint)
	}

	if len(a.Organizations) > 0 {
		stream.WriteMore()
		stream.WriteObjectField("organizations")
		stream.WriteArrayStart()

		for i, org := range a.Organizations {
			if i > 0 {
				stream.WriteMore()
			}

			stream.WriteString(org)
		}

		stream.WriteArrayEnd()
	}

	if len(a.RelatedResources) > 0 {
		stream.WriteMore()
		stream.WriteObjectField("related_resources")
		stream.WriteArrayStart()

		for i, res := range a.RelatedResources {
			if i > 0 {
				stream.WriteMore()
			}

			stream.WriteVal(res)
		}

		stream.WriteArrayEnd()
	}

	if len(a.Authors) > 0 {
		stream.WriteMore()
		stream.WriteObjectField("authors")
		stream.WriteArrayStart()

		for i, author := range a.Authors {
			if i > 0 {
				stream.WriteMore()
			}

			stream.WriteVal(author)
		}

		stream.WriteArrayEnd()
	}

	if len(a.Schemas) > 0 {
		stream.WriteMore()
		stream.WriteObjectField("schemas")
		stream.WriteArrayStart()

		for i, schema := range a.Schemas {
			if i > 0 {
				stream.WriteMore()
			}

			stream.WriteVal(schema)
		}

		stream.WriteArrayEnd()
	}

	if len(a.Custom) > 0 {
		stream.WriteMore()
		stream.WriteObjectField("custom")
		stream.WriteObjectStart()

		for key, value := range a.Custom {
			stream.WriteObjectField(key)
			stream.WriteVal(value)
		}

		stream.WriteObjectEnd()
	}

	stream.WriteObjectEnd()
}
