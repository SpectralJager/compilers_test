package ir

import (
	"fmt"
	"grimlang/internal/object"
	tp "grimlang/internal/type"
	"log"
	"strings"
)

type MetaData struct {
	Type     *tp.Type
	ArgTypes []tp.Type
}

type Module struct {
	Name      string
	Functions []Function
	Global    Function
	Meta      MetaData
}

func NewModule(name string) Module {
	return Module{
		Name: name,
	}
}

func (m *Module) SetGlobal(global Function) {
	m.Global = global
}

func (m *Module) PushFunctions(functions ...Function) {
	m.Functions = append(m.Functions, functions...)
}

func (m Module) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "--- %s\n", m.Name)
	fmt.Fprintf(&buf, "%s", m.Global)
	for _, f := range m.Functions {
		fmt.Fprintf(&buf, "%s", f)
	}
	return buf.String()
}

type Function struct {
	Name string
	Body []Instruction
	Meta MetaData
}

func NewFunction(name string) Function {
	return Function{
		Name: name,
	}
}

func (f *Function) PushInstructions(instructions ...Instruction) {
	f.Body = append(f.Body, instructions...)
}

func (f *Function) InsertInstructions(index int, instructions ...Instruction) error {
	if index < 0 || index > len(f.Body) {
		return fmt.Errorf("can't insert instructions, index %d out of bounds", index)
	}
	f.Body = append(f.Body[:index], append(instructions, f.Body[index:]...)...)
	return nil
}

func (f Function) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, ".%s =>\n", f.Name)
	fmt.Fprint(&buf, ";")
	if f.Meta.ArgTypes != nil {
		for _, arg := range f.Meta.ArgTypes {
			fmt.Fprintf(&buf, " %s", arg)
		}
		fmt.Fprint(&buf, " ->")
		if f.Meta.Type != nil {
			fmt.Fprintf(&buf, " %s", f.Meta.Type)
		} else {
			fmt.Fprint(&buf, " void")
		}
	} else {
		fmt.Fprint(&buf, " void ->")
		if f.Meta.Type != nil {
			fmt.Fprintf(&buf, " %s", f.Meta.Type)
		} else {
			fmt.Fprint(&buf, " void")
		}
	}
	fmt.Fprint(&buf, "\n")
	for _, code := range f.Body {
		fmt.Fprintf(&buf, "\t%s\n", code)
	}
	return buf.String()
}

type Instruction struct {
	Kind             InstructionKind
	Identifier       string
	Type             tp.Type
	Value            object.Object
	Label, LabelElse string
	Meta             MetaData
}

func VarNew(ident string, dataType tp.Type) Instruction {
	return Instruction{Kind: OP_VAR_NEW, Identifier: ident, Type: dataType}
}
func VarFree(ident string) Instruction {
	return Instruction{Kind: OP_VAR_FREE, Identifier: ident}
}
func VarSave(ident string) Instruction {
	return Instruction{Kind: OP_VAR_SAVE, Identifier: ident}
}
func VarLoad(ident string) Instruction {
	return Instruction{Kind: OP_VAR_LOAD, Identifier: ident}
}

func ConstLoad(val object.Object) Instruction {
	return Instruction{Kind: OP_CONST_LOAD, Value: val}
}

func Call(ident string, count object.Object) Instruction {
	if count.Kind != object.Int {
		log.Fatalf("invalid call 'count' argument: %T, expected Int", count)
	}
	return Instruction{Kind: OP_CALL, Identifier: ident, Value: count}
}
func Return(count object.Object) Instruction {
	if count.Kind != object.Int {
		log.Fatalf("invalid return 'count' argument: %T, expected Int", count)
	}
	return Instruction{Kind: OP_RETURN, Value: count}
}

func Label(label string) Instruction {
	return Instruction{Kind: OP_LABEL, Label: label}
}

func Br(label string) Instruction {
	return Instruction{Kind: OP_BR, Label: label}
}
func BrTrue(labelThen, labelElse string) Instruction {
	return Instruction{Kind: OP_BR_TRUE, Label: labelThen, LabelElse: labelElse}
}
func BrFalse(labelThen, labelElse string) Instruction {
	return Instruction{Kind: OP_BR_FALSE, Label: labelThen, LabelElse: labelElse}
}

func StackType(dataType tp.Type) Instruction {
	return Instruction{Kind: OP_STACK_TYPE, Type: dataType}
}

func (instr Instruction) GetMeta() MetaData {
	return instr.Meta
}

func (instr *Instruction) PutMeta(meta MetaData) {
	instr.Meta = meta
}

func (instr Instruction) String() string {
	switch instr.Kind {
	case OP_NOP:
		return string(instr.Kind)
	case OP_VAR_NEW:
		return fmt.Sprintf("%s: %s -> %s", instr.Kind, instr.Identifier, instr.Type)
	case OP_VAR_FREE:
		return fmt.Sprintf("%s: %s", instr.Kind, instr.Identifier)
	case OP_VAR_LOAD:
		return fmt.Sprintf("%s: %s", instr.Kind, instr.Identifier)
	case OP_VAR_SAVE:
		return fmt.Sprintf("%s: %s", instr.Kind, instr.Identifier)
	case OP_CONST_LOAD:
		return fmt.Sprintf("%s: #%s", instr.Kind, instr.Value)
	case OP_CALL:
		return fmt.Sprintf("%s: %s #%s", instr.Kind, instr.Identifier, instr.Value)
	case OP_RETURN:
		return fmt.Sprintf("%s: #%s", instr.Kind, instr.Value)
	case OP_LABEL:
		return fmt.Sprintf("%s: %s", instr.Kind, instr.Label)
	case OP_BR:
		return fmt.Sprintf("%s: %s", instr.Kind, instr.Label)
	case OP_BR_TRUE:
		return fmt.Sprintf("%s: %s ?: %s", instr.Kind, instr.Label, instr.LabelElse)
	case OP_BR_FALSE:
		return fmt.Sprintf("%s: %s ?: %s", instr.Kind, instr.Label, instr.LabelElse)
	case OP_STACK_TYPE:
		return fmt.Sprintf("%s: %s", instr.Kind, instr.Type)
	default:
		return "unknown"
	}
}

type InstructionKind string

const (
	OP_NOP InstructionKind = "nop"

	OP_VAR_NEW  InstructionKind = "var.new"
	OP_VAR_FREE InstructionKind = "var.free"
	OP_VAR_LOAD InstructionKind = "var.load"
	OP_VAR_SAVE InstructionKind = "var.save"

	OP_CONST_LOAD InstructionKind = "const.load"

	OP_CALL   InstructionKind = "call"
	OP_RETURN InstructionKind = "ret"

	OP_LABEL InstructionKind = "label"

	OP_BR       InstructionKind = "br"
	OP_BR_TRUE  InstructionKind = "br.true"
	OP_BR_FALSE InstructionKind = "br.false"

	OP_STACK_TYPE InstructionKind = "stack.type"
)
