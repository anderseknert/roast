package encoding

import (
	"testing"

	"github.com/open-policy-agent/opa/ast"
)

func TestMarshalBody(t *testing.T) {
	t.Parallel()

	body := &ast.Body{}

	_, err := JSON().Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
}
