package runtime

type Enviroment interface {
	Name() string
	Parent() Enviroment
	SearchLocal(string) Symbol
	Search(string) Symbol
	Insert(Symbol)
}

type env struct {
	name    string
	parent  Enviroment
	symbols map[string]Symbol
}

func (env *env) Name() string {
	return env.name
}

func (env *env) Parent() Enviroment {
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
		panic("symbol " + symbol.Name + " already defined in " + env.name)
	}
	env.symbols[symbol.Name()] = symbol
}
