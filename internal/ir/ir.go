package ir

import (
	"bytes"
	"fmt"
)

/*
package: *Name*
Constants: // only int, float, string
	0: 12
	1: "some string"
	2: 12.2
	3: 3.14
	...
Globals:
	a int
	b string
	c float
	d bool
	sum func[int int]<int>
	main func[void]<void>
init:
00000000|> global_set a $0;
00000001|> global_set b $1;
00000002|> global_set c $2;
00000003|> global_load b;
00000004|> call_builtin len;
00000005|> global_load a;
00000006|> call_builtin neq;
00000007|> global_save d;
functions:
=== main:
Arguments:
	void
Locals:
	result int
Body:
00000000|> load $12;
00000001|> load $11;
00000002|> call sum;
00000003|> local_save result;
00000004|> local_load result;
00000005|> call_builtin print;
00000006|> ret;
=== main:
Arguments:
	arg1 int
	arg2 int
Locals:
	void
Body:
00000000|> call_builtin iadd;
00000001|> ret;
*/

type IR interface {
	fmt.Stringer
	ir()
}

type IConstant interface {
	IR
	constIR()
}

type ISymbolDef interface {
	IR
	symdefIR()
}

type IInstruction interface {
	IR
	instrIR()
}

type IDataType interface {
	IR
	dtIR()
}

type Program struct {
	Name      string
	Constants []IConstant
	Globals   map[string]ISymbolDef
	InitCode  []IInstruction
	Functions map[string]Function
}

type Function struct {
	Name     string
	Locals   map[string]ISymbolDef
	BodyCode []IInstruction
}

// Constants
func (*Integer) constIR() {}
func (*Float) constIR()   {}
func (*String) constIR()  {}

type Integer struct {
	Value int
}

type Float struct {
	Value float64
}

type String struct {
	Value string
}

// Defenitions
func (*VaribleDef) symdefIR()  {}
func (*FunctionDef) symdefIR() {}

type VaribleDef struct {
	Name string
	Type IDataType
}

type FunctionDef struct {
	Name      string
	Arguments []IDataType
	Returns   []IDataType
}

// DataTypes
func (*Primitive) dtIR() {}

type Primitive struct {
	Name string
}

// Instructions
func (*GlobalSet) instrIR()       {}
func (*GlobalSave) instrIR()      {}
func (*GlobalLoad) instrIR()      {}
func (*LocalSet) instrIR()        {}
func (*LocalSave) instrIR()       {}
func (*LocalLoad) instrIR()       {}
func (*Load) instrIR()            {}
func (*Jump) instrIR()            {}
func (*ConditionalJump) instrIR() {}
func (*RelativeJump) instrIR()    {}
func (*CallBuildin) instrIR()     {}
func (*Call) instrIR()            {}
func (*Return) instrIR()          {}

type GlobalSet struct {
	Symbol     string
	ConstIndex int
}

type GlobalSave struct {
	Symbol string
}

type GlobalLoad struct {
	Symbol string
}

type LocalSet struct {
	Symbol     string
	ConstIndex int
}

type LocalSave struct {
	Symbol string
}

type LocalLoad struct {
	Symbol string
}

type Load struct {
	ConstIndex int
}

type Jump struct {
	Address int
}

type ConditionalJump struct {
	Address int
}

type RelativeJump struct {
	Count int
}

type Call struct {
	FuncName string
}

type CallBuildin struct {
	FuncName string
}

type Return struct {
	Count int
}

// Stringers
func (s *Program) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package: %s\n", s.Name)
	fmt.Fprint(&buf, "constants: \n")
	for i, c := range s.Constants {
		fmt.Fprintf(&buf, "\t%d: %s\n", i, c.String())
	}
	fmt.Fprint(&buf, "globals: \n")
	for _, g := range s.Globals {
		fmt.Fprintf(&buf, "\t%s\n", g.String())
	}
	fmt.Fprint(&buf, "init: \n")
	for i, code := range s.InitCode {
		fmt.Fprintf(&buf, "%08x|> %s\n", i, code.String())
	}
	fmt.Fprint(&buf, "functions: \n")
	for _, function := range s.Functions {
		fmt.Fprint(&buf, function.String())
	}
	return buf.String()
}

func (s *Function) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "=== %s:\n", s.Name)
	fmt.Fprint(&buf, "locals: \n")
	for _, loc := range s.Locals {
		fmt.Fprintf(&buf, "\t%s\n", loc.String())
	}
	fmt.Fprint(&buf, "body: \n")
	for i, code := range s.BodyCode {
		fmt.Fprintf(&buf, "%08x|> %s\n", i, code.String())
	}
	return buf.String()
}

func (s *Integer) String() string {
	return fmt.Sprintf("%d", s.Value)
}

func (s *Float) String() string {
	return fmt.Sprintf("%f", s.Value)
}

func (s *String) String() string {
	return s.Value
}

func (s *VaribleDef) String() string {
	return fmt.Sprintf("%s %s", s.Name, s.Type.String())
}

func (s *FunctionDef) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s func[", s.Name)
	for i, t := range s.Arguments {
		if i == len(s.Arguments)-1 {
			fmt.Fprintf(&buf, "%s", t.String())
			break
		}
		fmt.Fprintf(&buf, "%s ", t.String())
	}
	fmt.Fprint(&buf, "]<")
	for i, t := range s.Returns {
		if i == len(s.Returns)-1 {
			fmt.Fprintf(&buf, "%s", t.String())
			break
		}
		fmt.Fprintf(&buf, "%s ", t.String())
	}
	fmt.Fprint(&buf, ">")
	return buf.String()
}

func (s *Primitive) String() string {
	return s.Name
}

func (s *GlobalSet) String() string {
	return fmt.Sprintf("global_set %s $%d;", s.Symbol, s.ConstIndex)
}

func (s *GlobalLoad) String() string {
	return fmt.Sprintf("global_load %s;", s.Symbol)
}

func (s *GlobalSave) String() string {
	return fmt.Sprintf("global_save %s;", s.Symbol)
}

func (s *LocalSet) String() string {
	return fmt.Sprintf("local_set %s $%d;", s.Symbol, s.ConstIndex)
}

func (s *LocalLoad) String() string {
	return fmt.Sprintf("local_load %s;", s.Symbol)
}

func (s *LocalSave) String() string {
	return fmt.Sprintf("local_save %s;", s.Symbol)
}

func (s *Load) String() string {
	return fmt.Sprintf("load $%d;", s.ConstIndex)
}

func (s *Jump) String() string {
	return fmt.Sprintf("jump %08x;", s.Address)
}

func (s *ConditionalJump) String() string {
	return fmt.Sprintf("cond_jump %08x;", s.Address)
}

func (s *RelativeJump) String() string {
	return fmt.Sprintf("rel_jump %d;", s.Count)
}

func (s *CallBuildin) String() string {
	return fmt.Sprintf("call_buildin %s;", s.FuncName)
}

func (s *Call) String() string {
	return fmt.Sprintf("call %s;", s.FuncName)
}

func (s *Return) String() string {
	return fmt.Sprintf("return %d;", s.Count)
}

// irs
func (*Program) ir()         {}
func (*Function) ir()        {}
func (*Integer) ir()         {}
func (*Float) ir()           {}
func (*String) ir()          {}
func (*Primitive) ir()       {}
func (*VaribleDef) ir()      {}
func (*FunctionDef) ir()     {}
func (*GlobalSet) ir()       {}
func (*GlobalLoad) ir()      {}
func (*GlobalSave) ir()      {}
func (*LocalSet) ir()        {}
func (*LocalLoad) ir()       {}
func (*LocalSave) ir()       {}
func (*Load) ir()            {}
func (*CallBuildin) ir()     {}
func (*Call) ir()            {}
func (*Jump) ir()            {}
func (*ConditionalJump) ir() {}
func (*RelativeJump) ir()    {}
func (*Return) ir()          {}
