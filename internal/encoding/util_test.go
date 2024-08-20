package encoding

import (
	"github.com/open-policy-agent/opa/ast"
)

func MustParseModule(policy string) *ast.Module {
	return ast.MustParseModuleWithOpts(policy, ast.ParserOptions{ProcessAnnotation: true})
}
