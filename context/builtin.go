package context

import (
	"grimlang/builtin"
	"grimlang/dtype"
	"grimlang/symbol"
)

func NewBuiltinContext() Context {
	return &_context{
		scope: "builtin",
		symbols: []symbol.Symbol{
			&symbol.BuiltinFunctionSymbol{
				Identifier: "exit",
				Type: &dtype.FunctionType{
					Args: []dtype.Type{
						&dtype.IntType{},
					},
				},
				Callee: builtin.Exit,
			},
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
					&symbol.BuiltinFunctionSymbol{
						Identifier: "sub",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.IntType{},
								},
							},
							Return: &dtype.IntType{},
						},
						Callee: builtin.IntSub,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "mul",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.IntType{},
								},
							},
							Return: &dtype.IntType{},
						},
						Callee: builtin.IntMul,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "div",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.IntType{},
								},
							},
							Return: &dtype.IntType{},
						},
						Callee: builtin.IntDiv,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "lt",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.IntType{},
								},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.IntLt,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "gt",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.IntType{},
								},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.IntGt,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "leq",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.IntType{},
								},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.IntLeq,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "geq",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.IntType{},
								},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.IntGeq,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "eq",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.IntType{},
								},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.IntEq,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "toString",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.IntType{},
							},
							Return: &dtype.StringType{},
						},
						Callee: builtin.IntToString,
					},
				},
			},
			&symbol.ModuleSymbol{
				Identifier: "io",
				Symbols: []symbol.Symbol{
					&symbol.BuiltinFunctionSymbol{
						Identifier: "println",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.StringType{},
								},
							},
						},
						Callee: builtin.IoPrintln,
					},
				},
			},
		},
	}
}
