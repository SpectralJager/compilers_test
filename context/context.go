package context

import (
	"fmt"
	"grimlang/symbol"
)

type Context interface {
	Scope() string
	Search(string) symbol.Symbol
	Insert(symbol.Symbol) error
}

func NewContext(scope string, parent Context) Context {
	return &_context{
		scope:  scope,
		parent: parent,
	}
}

type _context struct {
	scope   string
	parent  Context
	symbols []symbol.Symbol
}

func (ctx *_context) Scope() string {
	return ctx.scope
}

func (ctx *_context) Search(ident string) symbol.Symbol {
	for _, sym := range ctx.symbols {
		if sym.Name() == ident {
			return sym
		}
	}
	if ctx.parent != nil {
		return ctx.parent.Search(ident)
	}
	return nil
}

func (ctx *_context) Insert(sym symbol.Symbol) error {
	if res := ctx.Search(sym.Name()); res != nil {
		return fmt.Errorf("symbol '%s' already defined in scope '%s'", sym.Name(), ctx.Scope())
	}
	ctx.symbols = append(ctx.symbols, sym)
	return nil
}
