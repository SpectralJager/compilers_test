package eval

import (
	"fmt"
	"grimlang/internal/core/frontend/ast"
	"grimlang/internal/core/frontend/tokens"
)

func Eval(node ast.Node, env *Env) (interface{}, error) {
	switch node := node.(type) {
	case *ast.PrefixExpr:
		return evalPrefix(*node, env)
	case *ast.Number, *ast.Float, *ast.String:
		return evalAtom(node.(ast.Atom))
	case *ast.DefExpr:
		return evalDef(*node, env)
	case *ast.SymbolExpr:
		return evalSymbol(*node, env)
	default:
		return nil, fmt.Errorf("unknow Node type: %+v", node)
	}
}

func evalDef(node ast.DefExpr, env *Env) (interface{}, error) {
	res, err := Eval(node.Value, env)
	if err != nil {
		return "", err
	}
	(*env)[node.Symbol.Value] = res
	return node.Symbol.Value, nil
}

func evalSymbol(node ast.SymbolExpr, env *Env) (interface{}, error) {
	if val, ok := (*env)[node.Symbol.Value]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("undefined symbol %v", &node.Symbol.Value)
}

func evalPrefix(node ast.PrefixExpr, env *Env) (interface{}, error) {
	values := make([]interface{}, len(node.Args))
	for i, argument := range node.Args {
		res, err := Eval(argument, env)
		if err != nil {
			return nil, err
		}
		values[i] = res
	}

	switch node.Operator.Type {
	case tokens.Add:
		res := 0.0
		for _, val := range values {
			temp := val
			res += temp
		}
		return res, nil
	case tokens.Sub:
		res := 0
		for i, val := range values {
			if i == 0 {
				res = val.(int)
				continue
			}
			res -= val.(int)
		}
		return res, nil
	case tokens.Mul:
		res := 0
		for i, val := range values {
			if i == 0 {
				res = val.(int)
				continue
			}
			res *= val.(int)
		}
		return res, nil
	case tokens.Div:
		res := 0
		for i, val := range values {
			if i == 0 {
				res = val.(int)
				continue
			}
			res /= val.(int)
		}
		return res, nil
	default:
		return nil, fmt.Errorf("wrong prefix expression symbol")
	}
}

func evalAtom(node ast.Atom) (interface{}, error) {
	return node.Value()
}
