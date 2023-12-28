package context

import "grimlang/symbol"

func NewBuiltinContext() Context {
	return &_context{
		scope:   "builtin",
		symbols: []symbol.Symbol{},
	}
}
