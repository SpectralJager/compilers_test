package compiler

import (
	"bytes"
	"fmt"
	"grimlang/internal/core/frontend/ast"
)

func Compile(node ast.Node, buf *bytes.Buffer) string {
	switch node := node.(type) {
	case *ast.DefSF:
		buf.WriteString(fmt.Sprintf("%s := ", node.Symb.Token.Value))
		Compile(node.Value, buf)
		buf.WriteString("\n")
	case *ast.RetSF:
		buf.WriteString("return ")
		Compile(node.Value, buf)
		buf.WriteString("\n")
	case *ast.Number:
		buf.WriteString(node.Token.Value)
	case *ast.String:
		buf.WriteString("\"" + node.Token.Value + "\"")
	case *ast.Bool:
		buf.WriteString(node.Token.Value)
	case *ast.Symbol:
		buf.WriteString(node.Token.Value)
	case *ast.FnSF:
		buf.WriteString(fmt.Sprintf("func %s (", node.Symb.Token.Value))
		for _, v := range node.Args {
			Compile(&v, buf)
			buf.WriteString(" any,")
		}
		buf.WriteString(") any{\n")
		for _, v := range node.Body {
			Compile(v, buf)
		}
		buf.WriteString("}\n")
	case *ast.SymbolExpr:
		buf.WriteString(fmt.Sprintf("%s(", node.Symb.Token.Value))
		for _, v := range node.Args {
			Compile(v, buf)
			buf.WriteString(",")
		}
		buf.WriteString(")")
	default:
		panic("unexpected node type" + node.String())
	}
	return ""
}
