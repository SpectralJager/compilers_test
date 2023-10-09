package ir

import (
	"fmt"
	"log"
	"strings"
)

type IR interface {
	fmt.Stringer
}

type ModuleIR struct {
	Name      string
	Init      []InstrIR
	Functions map[string]FunctionIR
}

func NewModule(name string) ModuleIR {
	return ModuleIR{
		Name:      name,
		Functions: make(map[string]FunctionIR),
		Init:      make([]InstrIR, 0),
	}
}

func (m *ModuleIR) WriteInstrs(instr ...InstrIR) {
	m.Init = append(m.Init, instr...)
}

func (m *ModuleIR) WriteFunctions(fns ...FunctionIR) {
	for _, f := range fns {
		if _, ok := m.Functions[f.Name.String()]; ok {
			log.Fatalf("function %s already exists", f.Name)
		}
		m.Functions[f.Name.String()] = f
	}
}

func (ir *ModuleIR) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "=== %s\n", ir.Name)
	for x, instr := range ir.Init {
		fmt.Fprintf(&buf, "%08x %s\n", x, instr.String())
	}
	for _, fn := range ir.Functions {
		fmt.Fprintf(&buf, "%s\n", fn.String())
	}
	return buf.String()
}

type FunctionIR struct {
	Name SymbolIR
	Code []InstrIR
}

func NewFunction(name SymbolIR) FunctionIR {
	return FunctionIR{Name: name}
}

func (ir *FunctionIR) WriteInstrs(instrs ...InstrIR) {
	ir.Code = append(ir.Code, instrs...)
}

func (ir *FunctionIR) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, ".%s:\n", ir.Name.String())
	for x, instr := range ir.Code {
		fmt.Fprintf(&buf, "%08x %s\n", x, instr.String())
	}
	return buf.String()
}

type IntIR struct {
	Value int
}

func NewInt(val int) IntIR {
	return IntIR{Value: val}
}

func (ir *IntIR) String() string {
	return fmt.Sprint(ir.Value)
}

type SymbolIR struct {
	Primary   string
	Secondary *SymbolIR
}

func NewSymbol(p string, s *SymbolIR) SymbolIR {
	return SymbolIR{
		Primary:   p,
		Secondary: s,
	}
}

func (ir *SymbolIR) String() string {
	if ir.Secondary == nil {
		return ir.Primary
	}
	return fmt.Sprintf("%s/%s", ir.Primary, ir.Secondary.String())
}

type TypeIR struct {
	Primary SymbolIR
	Generic []TypeIR
}

func NewType(p SymbolIR, g ...TypeIR) TypeIR {
	return TypeIR{
		Primary: p,
		Generic: g,
	}
}

func (ir *TypeIR) String() string {
	if ir.Generic == nil {
		return ir.Primary.String()
	}
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s<", ir.Primary.String())
	for _, t := range ir.Generic {
		fmt.Fprintf(&buf, "%s,", t.String())
	}
	fmt.Fprint(&buf, ">")
	return buf.String()
}
