package context

import (
	"fmt"
	"grimlang/builtin"
	builtinFloat "grimlang/builtin/float"
	builtinInt "grimlang/builtin/int"
	builtinIo "grimlang/builtin/io"
	builtinList "grimlang/builtin/list"
	"grimlang/builtin/math"
	builtinString "grimlang/builtin/string"
	"grimlang/object"
)

type Context interface {
	Scope() string
	Search(string) object.Symbol
	Insert(object.Symbol) error
}

func NewContext(scope string, prev Context) Context {
	return &_context{
		scope:   scope,
		symbols: []object.Symbol{},
		prev:    prev,
	}
}

type _context struct {
	scope   string
	symbols []object.Symbol
	prev    Context
}

func (ctx *_context) Scope() string {
	return ctx.scope
}
func (ctx *_context) Search(ident string) object.Symbol {
	for _, symb := range ctx.symbols {
		if symb.Name() == ident {
			return symb
		}
	}
	if ctx.prev == nil {
		return nil
	}
	return ctx.prev.Search(ident)
}
func (ctx *_context) Insert(symbol object.Symbol) error {
	for _, symb := range ctx.symbols {
		if symb.Name() == symbol.Name() {
			return fmt.Errorf("symbols %s already defined in %s", symbol.Name(), ctx.scope)
		}
	}
	ctx.symbols = append(ctx.symbols, symbol)
	return nil
}

func NewBuiltinContext() Context {
	return &_context{
		scope: "builtin",
		symbols: []object.Symbol{
			&object.SymbolBuiltin{
				Identifier: "exit",
				FunctionType: object.DTypeFunction{
					ArgumentsType: []object.DType{
						&object.DTypeInt{},
					},
				},
				Fn: builtin.Exit,
			},
			&object.SymbolModule{
				Identifier: "int",
				Symbols: []object.Symbol{
					&object.SymbolBuiltin{
						Identifier: "add",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeVariatic{
									ChildType: &object.DTypeInt{},
								},
							},
							ReturnType: &object.DTypeInt{},
						},
						Fn: builtinInt.IntAdd,
					},
					&object.SymbolBuiltin{
						Identifier: "sub",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeVariatic{
									ChildType: &object.DTypeInt{},
								},
							},
							ReturnType: &object.DTypeInt{},
						},
						Fn: builtinInt.IntSub,
					},
					&object.SymbolBuiltin{
						Identifier: "lt",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeVariatic{
									ChildType: &object.DTypeInt{},
								},
							},
							ReturnType: &object.DTypeBool{},
						},
						Fn: builtinInt.IntLt,
					},
					&object.SymbolBuiltin{
						Identifier: "toString",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeInt{},
							},
							ReturnType: &object.DTypeString{},
						},
						Fn: builtinInt.IntToString,
					},
				},
			},
			&object.SymbolModule{
				Identifier: "float",
				Symbols: []object.Symbol{
					&object.SymbolBuiltin{
						Identifier: "add",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeVariatic{
									ChildType: &object.DTypeFloat{},
								},
							},
							ReturnType: &object.DTypeFloat{},
						},
						Fn: builtinFloat.FloatAdd,
					},
					&object.SymbolBuiltin{
						Identifier: "sub",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeVariatic{
									ChildType: &object.DTypeFloat{},
								},
							},
							ReturnType: &object.DTypeFloat{},
						},
						Fn: builtinFloat.FloatSub,
					},
					&object.SymbolBuiltin{
						Identifier: "toString",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeFloat{},
							},
							ReturnType: &object.DTypeString{},
						},
						Fn: builtinFloat.FloatToString,
					},
				},
			},
			&object.SymbolModule{
				Identifier: "string",
				Symbols: []object.Symbol{
					&object.SymbolBuiltin{
						Identifier: "format",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeString{},
								&object.DTypeVariatic{
									ChildType: &object.DTypeString{},
								},
							},
							ReturnType: &object.DTypeString{},
						},
						Fn: builtinString.Format,
					},
				},
			},
			&object.SymbolModule{
				Identifier: "list",
				Symbols: []object.Symbol{
					&object.SymbolBuiltin{
						Identifier: "len",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeList{},
							},
							ReturnType: &object.DTypeInt{},
						},
						Fn: builtinList.Len,
					},
					&object.SymbolBuiltin{
						Identifier: "get",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeList{},
								&object.DTypeInt{},
							},
							ReturnType: &object.DTypeAny{},
						},
						Fn: builtinList.Get,
					},
				},
			},
			&object.SymbolModule{
				Identifier: "io",
				Symbols: []object.Symbol{
					&object.SymbolBuiltin{
						Identifier: "println",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeVariatic{
									ChildType: &object.DTypeString{},
								},
							},
						},
						Fn: builtinIo.Println,
					},
				},
			},
			&object.SymbolModule{
				Identifier: "math",
				Symbols: []object.Symbol{
					&object.SymbolBuiltin{
						Identifier: "sqrt",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeFloat{},
							},
							ReturnType: &object.DTypeFloat{},
						},
						Fn: math.Sqrt,
					},
					&object.SymbolBuiltin{
						Identifier: "pow",
						FunctionType: object.DTypeFunction{
							ArgumentsType: []object.DType{
								&object.DTypeFloat{},
								&object.DTypeFloat{},
							},
							ReturnType: &object.DTypeFloat{},
						},
						Fn: math.Pow,
					},
				},
			},
		},
	}
}
