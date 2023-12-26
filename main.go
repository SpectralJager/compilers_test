package main

import "grimlang/ast"

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
