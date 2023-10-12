package analysis

import (
	"fmt"
	tp "grimlang/internal/type"
	"strings"
)

type SymbolTable []Symbol

func (t *SymbolTable) Add(sm Symbol) error {
	for _, s := range *t {
		if s.Name() == sm.Name() {
			if s.Kind() == "variable" && sm.Kind() == "variable" {
				if s.String() == sm.String() {
					goto exists
				}
				goto ok
			}
			goto exists
		}
	}
ok:
	*t = append(*t, sm)
	return nil
exists:
	return fmt.Errorf("symbol %s already exists", sm.Name())
}

func (t SymbolTable) GetVars(scope string) []Symbol {
	ret := make([]Symbol, 0)
	for _, s := range t {
		switch s := s.(type) {
		case *VariableSymbol:
			if s.Scope == scope {
				ret = append(ret, s)
			}
		default:
			fmt.Printf("%#v: unknown symbol", s)
		}
	}
	return ret
}

type Symbol interface {
	fmt.Stringer
	Name() string
	Kind() string
}

type VariableSymbol struct {
	Ident string
	Type  tp.Type
	Scope string
}

func (v *VariableSymbol) String() string {
	if v.Scope != "" {
		return fmt.Sprintf("%s:%s of %s", v.Ident, v.Type, v.Scope)
	}
	return fmt.Sprintf("%s:%s", v.Ident, v.Type)
}

func (v *VariableSymbol) Kind() string {
	return "variable"
}

func (v *VariableSymbol) Name() string {
	return v.Ident
}

type FunctionSymbol struct {
	Ident       string
	Args        []tp.Type
	ReturnTypes []tp.Type
}

func (f *FunctionSymbol) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s[", f.Ident)
	if f.Args != nil {
		for _, v := range f.Args {
			fmt.Fprintf(&buf, " %s", v.String())
		}
	} else {
		fmt.Fprintf(&buf, "void")
	}
	fmt.Fprintf(&buf, "] -> ")
	if f.ReturnTypes != nil {
		for _, v := range f.ReturnTypes {
			fmt.Fprintf(&buf, "%s, ", v.String())
		}
	} else {
		fmt.Fprintf(&buf, "void")
	}
	return buf.String()
}

func (f *FunctionSymbol) Kind() string {
	return "function"
}

func (f *FunctionSymbol) Name() string {
	return f.Ident
}
