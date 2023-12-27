package main

import (
	"grimlang/ast"
	"grimlang/builtin"
	"grimlang/context"
	"grimlang/dtype"
	"grimlang/eval"
	"grimlang/symbol"
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
			Identifier: "add",
			Arguments: []*ast.VariableDefn{
				{Identifier: "a", Type: &ast.IntType{}},
				{Identifier: "b", Type: &ast.IntType{}},
			},
			Return: &ast.IntType{},
			Body: []ast.Local{
				&ast.ReturnStmt{
					Value: &ast.SymbolCall{
						Call: &ast.SymbolExpr{
							Identifier: "int",
							Next:       &ast.SymbolExpr{Identifier: "add"},
						},
						Arguments: []ast.Expression{
							&ast.SymbolExpr{Identifier: "a"},
							&ast.SymbolExpr{Identifier: "b"},
						},
					},
				},
			},
		},
		&ast.FunctionDecl{
			Identifier: "main",
			Arguments:  nil,
			Return:     nil,
			Body: []ast.Local{
				&ast.SymbolCall{
					Call: &ast.SymbolExpr{Identifier: "exit"},
					Arguments: []ast.Expression{
						&ast.SymbolCall{
							Call: &ast.SymbolExpr{Identifier: "add"},
							Arguments: []ast.Expression{
								&ast.IntAtom{Value: 10},
								&ast.SymbolExpr{Identifier: "n", Next: nil},
							},
						},
					},
				},
			},
		},
	},
}

func main() {
	builtinContext := context.NewContext("builtint", nil)
	builtinContext.Insert(
		&symbol.BuiltinFunctionSymbol{
			Identifier: "exit",
			Type: &dtype.FunctionType{
				Args: []dtype.Type{
					&dtype.IntType{},
				},
				Return: nil,
			},
			Callee: builtin.Exit,
		},
	)
	builtinContext.Insert(
		&symbol.ModuleSymbol{
			Identifier: "int",
			Symbols: []symbol.Symbol{
				&symbol.BuiltinFunctionSymbol{
					Identifier: "add",
					Type: &dtype.FunctionType{
						Args: []dtype.Type{
							&dtype.VariaticType{
								Child: &dtype.IntType{},
							},
						},
						Return: &dtype.IntType{},
					},
					Callee: builtin.IntAdd,
				},
			},
		},
	)

	if err := new(eval.EvalState).EvalModule(builtinContext, tree); err != nil {
		log.Fatalf("something goes wrong -> %s", err.Error())
	}
	os.Exit(0)
}
