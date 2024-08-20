package encoding

import (
	"encoding/base64"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/open-policy-agent/opa/ast"
	"unsafe"
)

type locationCodec struct{}

func (*locationCodec) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (*locationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	loc := *((*ast.Location)(ptr))

	stream.WriteString(fmt.Sprintf("%d:%d:%s", loc.Row, loc.Col, base64.StdEncoding.EncodeToString(loc.Text)))
}
