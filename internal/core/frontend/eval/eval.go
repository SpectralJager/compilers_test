package eval

import (
	"fmt"
	"grimlang/internal/core/frontend/ast"
	"grimlang/internal/core/frontend/objects"
	"grimlang/internal/core/frontend/tokens"
	"strconv"
)

func Eval(node ast.Node, env *Env) (objects.Object, error) {
	switch node := node.(type) {
	case *ast.Number:
		return evalNumber(*node)
	default:
		return nil, fmt.Errorf("unknow Node type: %+v", node)
	}
}

func evalNumber(node ast.Number) (objects.Object, error) {
	switch node.Token.Type {
	case tokens.Int:
		res, err := strconv.Atoi(node.Token.Value)
		if err != nil {
			return nil, err
		}
		return &objects.ObjectInt{Value: res}, nil
	case tokens.Float:
		res, err := strconv.ParseFloat(node.Token.Value, 64)
		if err != nil {
			return nil, err
		}
		return &objects.ObjectFloat{Value: res}, nil
	default:
		return nil, fmt.Errorf("unsupported number type %s", node.Token.Type.String())
	}
}
