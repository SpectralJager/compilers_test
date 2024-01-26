package runtime

import (
	"fmt"
	"strings"
)

type Type interface {
	Kind() Kind
	Name() string
	String() string
	Compare(Type) bool
	Item() Type
	NumIns() int
	In(int) Type
	Out() Type
}

type typ struct {
	kind Kind
	item Type
	args []Type
	ret  Type
}

func (typ *typ) Kind() Kind {
	return typ.kind
}

func (typ *typ) Name() string {
	switch typ.kind {
	case TY_Void:
		return "void"
	case TY_Int:
		return "int"
	case TY_Float:
		return "float"
	case TY_Bool:
		return "bool"
	case TY_String:
		return "string"
	case TY_List:
		return "list"
	case TY_Function:
		return "fn"
	default:
		panic("can't get type name: unexpected kind of type")
	}
}

func (typ *typ) String() string {
	switch typ.kind {
	case TY_Int, TY_Float, TY_Bool, TY_String, TY_Void:
		return typ.Name()
	case TY_List:
		return fmt.Sprintf("%s<%s>", typ.Name(), typ.item.String())
	case TY_Function:
		args := []string{}
		for _, arg := range typ.args {
			args = append(args, arg.String())
		}
		return fmt.Sprintf("%s[%s]<%s>", typ.Name(), strings.Join(args, " "), typ.ret.String())
	default:
		panic("can't get string of type: unexpected kind of type")
	}
}

func (typ *typ) Compare(other Type) bool {
	if typ.kind != other.Kind() {
		return false
	}
	switch typ.kind {
	case TY_Int, TY_Float, TY_Bool, TY_String, TY_Void:
	case TY_List:
		return other.Item().Compare(typ.item)
	case TY_Function:
		if typ.NumIns() != other.NumIns() || !other.Out().Compare(typ.ret) {
			return false
		}
		for i, in := range typ.args {
			if !in.Compare(other.In(i)) {
				return false
			}
		}
	}
	return true
}

func (typ *typ) Item() Type {
	if typ.kind != TY_List {
		panic("can't get item type: type is not list")
	}
	if typ.item == nil {
		panic("can't get item type: item type is empty")
	}
	return typ.item
}

func (typ *typ) NumIns() int {
	if typ.kind != TY_Function {
		panic("can't get number of input arguments: type is not function")
	}
	return len(typ.args)
}

func (typ *typ) In(index int) Type {
	if typ.kind != TY_Function {
		panic("can't get type of input argument: type is not function")
	}
	if index >= typ.NumIns() && index < 0 {
		panic("can't get type of input argument: index out of bounds")
	}
	return typ.args[index]
}

func (typ *typ) Out() Type {
	if typ.kind != TY_Function {
		panic("can't get type of output: type is not function")
	}
	if typ.ret == nil {
		return NewVoidType()
	}
	return typ.ret
}

func NewVoidType() *typ {
	return &typ{
		kind: TY_Void,
	}
}

func NewIntType() *typ {
	return &typ{
		kind: TY_Int,
	}
}

func NewFloatType() *typ {
	return &typ{
		kind: TY_Float,
	}
}
func NewStringType() *typ {
	return &typ{
		kind: TY_String,
	}
}
func NewBoolType() *typ {
	return &typ{
		kind: TY_Bool,
	}
}

func NewListType(item Type) *typ {
	return &typ{
		kind: TY_List,
		item: item,
	}
}

func NewFunctionType(out Type, ins ...Type) *typ {
	return &typ{
		kind: TY_Function,
		args: ins,
		ret:  out,
	}
}
