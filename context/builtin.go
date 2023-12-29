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
					&symbol.BuiltinFunctionSymbol{
						Identifier: "toFloat",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.IntType{},
							},
							Return: &dtype.FloatType{},
						},
						Callee: builtin.IntToFloat,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "toBool",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.IntType{},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.IntToBool,
					},
				},
			},
			&symbol.ModuleSymbol{
				Identifier: "float",
				Symbols: []symbol.Symbol{
					&symbol.BuiltinFunctionSymbol{
						Identifier: "add",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.FloatType{},
								},
							},
							Return: &dtype.FloatType{},
						},
						Callee: builtin.FloatAdd,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "sub",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.FloatType{},
								},
							},
							Return: &dtype.FloatType{},
						},
						Callee: builtin.FloatSub,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "mul",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.FloatType{},
								},
							},
							Return: &dtype.FloatType{},
						},
						Callee: builtin.FloatMul,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "div",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.FloatType{},
								},
							},
							Return: &dtype.FloatType{},
						},
						Callee: builtin.FloatDiv,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "lt",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.FloatType{},
								},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.FloatLt,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "gt",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.FloatType{},
								},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.FloatGt,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "leq",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.FloatType{},
								},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.FloatLeq,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "geq",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.FloatType{},
								},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.FloatGeq,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "eq",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.FloatType{},
								},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.FloatEq,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "toString",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.FloatType{},
							},
							Return: &dtype.StringType{},
						},
						Callee: builtin.FloatToString,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "toInt",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.FloatType{},
							},
							Return: &dtype.IntType{},
						},
						Callee: builtin.FloatToInt,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "toBool",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.FloatType{},
							},
							Return: &dtype.BoolType{},
						},
						Callee: builtin.FloatToBool,
					},
				},
			},
			&symbol.ModuleSymbol{
				Identifier: "string",
				Symbols: []symbol.Symbol{
					&symbol.BuiltinFunctionSymbol{
						Identifier: "concat",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.VariaticType{
									Child: &dtype.StringType{},
								},
							},
							Return: &dtype.StringType{},
						},
						Callee: builtin.StringConcat,
					},
				},
			},
			&symbol.ModuleSymbol{
				Identifier: "list",
				Symbols: []symbol.Symbol{
					&symbol.BuiltinFunctionSymbol{
						Identifier: "len",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.ListType{},
							},
							Return: &dtype.IntType{},
						},
						Callee: builtin.ListLen,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "get",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.ListType{},
							},
							Return: &dtype.AnyType{},
						},
						Callee: builtin.ListGet,
					},
					&symbol.BuiltinFunctionSymbol{
						Identifier: "set",
						Type: &dtype.FunctionType{
							Args: []dtype.Type{
								&dtype.ListType{},
							},
						},
						Callee: builtin.ListSet,
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
