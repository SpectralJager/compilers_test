package eval

import (
	"fmt"
	"grimlang/internal/core/frontend/ast"
	"strconv"
)

func Eval(node ast.Node, env *Env) (interface{}, error) {
	switch node := node.(type) {
	case *ast.Int:
		return strconv.Atoi(node.Token.Value)
	default:
		return nil, fmt.Errorf("unknow Node type: %+v", node)
	}
}
