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

type IInstruction []byte

type IDataType interface {
	IR
	dtIR()
}

type Program struct {
	Name      string
	Constants []IConstant
	Globals   []ISymbolDef
	InitCode  *Code
	Functions map[string]*Function
}

type Function struct {
	Name     string
	Locals   []ISymbolDef
	BodyCode *Code
}

// Constants
func (*Integer) constIR() {}
func (*Float) constIR()   {}
func (*String) constIR()  {}
func (*True) constIR()    {}
func (*False) constIR()   {}

type Integer struct {
	Value int
}

type Float struct {
	Value float64
}

type String struct {
	Value string
}

type True struct{}
type False struct{}

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

// Stringers
func (s *Program) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package: %s\n", s.Name)
	fmt.Fprint(&buf, "constants: \n")
	for i, c := range s.Constants {
		fmt.Fprintf(&buf, "\t%d: %s\n", i, c.String())
	}
	fmt.Fprint(&buf, "globals: \n")
	for i, g := range s.Globals {
		fmt.Fprintf(&buf, "\t%d: %s\n", i, g.String())
	}
	fmt.Fprint(&buf, "init: \n")
	fmt.Fprint(&buf, s.InitCode.Disassembly())
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
	for i, loc := range s.Locals {
		fmt.Fprintf(&buf, "\t%d: %s\n", i, loc.String())
	}
	fmt.Fprint(&buf, "body: \n")
	fmt.Fprint(&buf, s.BodyCode.Disassembly())
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

func (*True) String() string {
	return "true"
}

func (*False) String() string {
	return "false"
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

// irs
func (*Program) ir()     {}
func (*Function) ir()    {}
func (*Integer) ir()     {}
func (*Float) ir()       {}
func (*String) ir()      {}
func (*True) ir()        {}
func (*False) ir()       {}
func (*Primitive) ir()   {}
func (*VaribleDef) ir()  {}
func (*FunctionDef) ir() {}
