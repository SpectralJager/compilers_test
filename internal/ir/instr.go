package ir

import (
	"fmt"
	"strings"
)

type InstrKind uint

const (
	OP_NOP InstrKind = iota

	// OP_GLOBAL_NEW
	// OP_GLOBAL_LOAD
	// OP_GLOBAL_SAVE

	// OP_LOCAL_NEW
	// OP_LOCAL_LOAD
	// OP_LOCAL_SAVE

	OP_VAR_NEW
	OP_VAR_LOAD
	OP_VAR_SAVE

	OP_CONST_LOAD

	OP_CALL
	OP_RETURN

	OP_LABEL

	OP_BR
	OP_BR_TRUE
	OP_BR_FALSE

	OP_STACK_POP
	OP_STACK_DUP
	OP_STACK_TYPE
)

func (ir InstrKind) String() string {
	switch ir {
	case OP_NOP:
		return "nop"
	case OP_VAR_NEW:
		return "var.new"
	case OP_VAR_LOAD:
		return "var.load"
	case OP_VAR_SAVE:
		return "var.save"
	case OP_CONST_LOAD:
		return "const.load"
	case OP_CALL:
		return "call"
	case OP_RETURN:
		return "return"
	case OP_LABEL:
		return "label"
	case OP_BR:
		return "br"
	case OP_BR_TRUE:
		return "br.true"
	case OP_BR_FALSE:
		return "br.false"
	case OP_STACK_POP:
		return "stack.pop"
	case OP_STACK_DUP:
		return "stack.dup"
	case OP_STACK_TYPE:
		return "stack.type"
	default:
		return fmt.Sprintf("undefined{%d}", ir)
	}
}

type InstrIR struct {
	Op   InstrKind
	Args []IR
}

func (ir *InstrIR) String() string {
	var buf strings.Builder
	if ir.Op == OP_LABEL {
		fmt.Fprintf(&buf, "%s:", ir.Args[0].String())
	} else {
		fmt.Fprintf(&buf, "%s", ir.Op.String())
		for _, arg := range ir.Args {
			fmt.Fprintf(&buf, " %s", arg.String())
		}
	}
	return buf.String()
}

func Nop() *InstrIR {
	return &InstrIR{Op: OP_NOP}
}

func VarNew(name *SymbolIR) *InstrIR {
	return &InstrIR{
		Op:   OP_VAR_NEW,
		Args: []IR{name},
	}
}

func VarLoad(name *SymbolIR) *InstrIR {
	return &InstrIR{
		Op:   OP_VAR_LOAD,
		Args: []IR{name},
	}
}

func VarSave(name *SymbolIR) *InstrIR {
	return &InstrIR{
		Op:   OP_VAR_SAVE,
		Args: []IR{name},
	}
}

func ConstLoadInt(val *IntIR) *InstrIR {
	return &InstrIR{
		Op:   OP_CONST_LOAD,
		Args: []IR{val},
	}
}

func Call(fn *SymbolIR, count *IntIR) *InstrIR {
	return &InstrIR{
		Op:   OP_CALL,
		Args: []IR{fn, count},
	}
}

func Lable(name *SymbolIR) *InstrIR {
	return &InstrIR{
		Op:   OP_LABEL,
		Args: []IR{name},
	}
}

func Br(label *SymbolIR) *InstrIR {
	return &InstrIR{
		Op:   OP_BR,
		Args: []IR{label},
	}
}

func BrTrue(thn, els *SymbolIR) *InstrIR {
	return &InstrIR{
		Op:   OP_BR_TRUE,
		Args: []IR{thn, els},
	}
}

func BrFalse(thn, els *SymbolIR) *InstrIR {
	return &InstrIR{
		Op:   OP_BR_FALSE,
		Args: []IR{thn, els},
	}
}

func StackPop() *InstrIR {
	return &InstrIR{Op: OP_STACK_POP}
}

func StackDup() *InstrIR {
	return &InstrIR{Op: OP_STACK_DUP}
}
func StackType(tp *TypeIR) *InstrIR {
	return &InstrIR{Op: OP_STACK_TYPE, Args: []IR{tp}}
}

func Return(count *IntIR) *InstrIR {
	return &InstrIR{Op: OP_RETURN, Args: []IR{count}}
}
