package context

import (
	"fmt"
	"grimlang/builtin"
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
		},
	}
}
