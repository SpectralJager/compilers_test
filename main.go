package main

import (
	"grimlang/ast"
	"grimlang/context"
	"grimlang/eval"
	"log"
	"os"
)

var tree = &ast.Module{
	Kind: ":main",
	Body: []ast.Global{
		&ast.ConstantDecl{
			Identifier: "n",
			Value:      &ast.IntAtom{Value: 20},
		},
		&ast.FunctionDecl{
			Identifier: "main",
			Arguments:  nil,
			Return:     nil,
			Body: []ast.Local{
				&ast.SymbolCall{
					Call: &ast.SymbolExpr{Identifier: "exit"},
					Arguments: []ast.Expression{
						&ast.SymbolExpr{Identifier: "n"},
					},
				},
			},
		},
	},
}

func main() {
	builtinContext := context.NewContext("builtint", nil)
	if err := eval.EvalModule(builtinContext, tree); err != nil {
		log.Fatalf("something goes wrong -> %s", err.Error())
	}
	os.Exit(0)
}
