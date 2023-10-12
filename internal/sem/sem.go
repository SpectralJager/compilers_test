package sem

import (
	"context"
	"grimlang/internal/ir"
	"log"
	"reflect"
)

type (
	_globals struct{}
)

func ModuleSemantic(m *ir.ModuleIR) {
	ctx := context.Background()
	globals := make([]*ir.SymbolIR, 0)
	for _, fns := range m.Functions {
		globals = appendSymbol(globals, &fns.Name)
	}

	ctx = context.WithValue(ctx, _globals{}, globals)
}

func appendSymbol(list []*ir.SymbolIR, sm *ir.SymbolIR) []*ir.SymbolIR {
	for _, n := range list {
		if reflect.DeepEqual(*n, *sm) {
			log.Fatalf("symbol %s already exists", n.String())
		}
	}
	list = append(list, sm)
	return list
}
