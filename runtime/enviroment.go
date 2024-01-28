package runtime

import (
	"fmt"
	"strings"
)

type Environment interface {
	String() string
	Name() string
	Parent() Environment
	SearchLocal(string) Symbol
	Search(string) Symbol
	Insert(Symbol)
}

type env struct {
	name    string
	parent  Environment
	symbols map[string]Symbol
}

func (env *env) Name() string {
	return env.name
}

func (env *env) Parent() Environment {
	return env.parent
}

func (env *env) SearchLocal(name string) Symbol {
	symbol, ok := env.symbols[name]
	if !ok {
		return nil
	}
	return symbol
}

func (env *env) Search(name string) Symbol {
	symbol := env.SearchLocal(name)
	if symbol == nil && env.parent != nil {
		return env.parent.Search(name)
	}
	return symbol
}

func (env *env) Insert(symbol Symbol) {
	if symbol := env.SearchLocal(symbol.Name()); symbol != nil {
		panic("symbol " + symbol.Name() + " already defined in " + env.name)
	}
	env.symbols[symbol.Name()] = symbol
}

func (env *env) String() string {
	symbols := []string{}
	for _, symb := range env.symbols {
		symbols = append(symbols, symb.String())
	}
	return fmt.Sprintf("%s |>\n\t%s\n", env.name, strings.Join(symbols, "\n\t"))
}

func NewEnvironment(name string, parent Environment) *env {
	return &env{
		name:    name,
		parent:  parent,
		symbols: map[string]Symbol{},
	}
}

func NewEnvironmentFromRecord(rec Litteral) *env {
	if rec.Kind() != LI_Record {
		panic("can't create Environment: rec should be record")
	}
	env := NewEnvironment(rec.Type().Name(), nil)
	for i := 0; i < rec.Type().NumFields(); i++ {
		fld := rec.Type().FieldByIndex(i)
		env.Insert(
			rec.Field(fld.Name()),
		)
	}
	return env
}
