package runtime

import (
	"fmt"
	"strings"
)

type TypeKind uint

const (
	TY_ANY TypeKind = iota
	TY_VOID
	TY_INT
	TY_FLOAT
	TY_STRING
	TY_BOOL
	TY_LIST
	TY_NULL
	TY_FUNC
	TY_RECORD
)

type Type struct {
	// general fields for all types
	Kind    TypeKind
	Name    string
	MdlPath string

	// child type for list, null
	Child *Type

	// arguments and return types for function
	Args   []*Type
	Return *Type

	// fields for record
	Fields []FieldType
}

func (tp *Type) String() string {
	switch tp.Kind {
	case TY_ANY, TY_VOID, TY_INT, TY_FLOAT, TY_BOOL, TY_STRING:
		return fmt.Sprintf("%s/%s", tp.MdlPath, tp.Name)
	case TY_NULL, TY_LIST:
		return fmt.Sprintf("%s/%s<%s>", tp.MdlPath, tp.Name, tp.Child.String())
	case TY_FUNC:
		args := []string{}
		for _, arg := range tp.Args {
			args = append(args, arg.String())
		}
		return fmt.Sprintf("%s/%s[%s]<%s>", tp.MdlPath, tp.Name, strings.Join(args, " "), tp.Return.String())
	case TY_RECORD:
		flds := []string{}
		for _, arg := range tp.Fields {
			flds = append(flds, arg.String())
		}
		return fmt.Sprintf("%s/%s{%s}", tp.MdlPath, tp.Name, strings.Join(flds, " "))
	}
	return ""
}

type FieldType struct {
	Name string
	Type *Type
}

func (fld *FieldType) String() string {
	return fmt.Sprintf("%s::%s", fld.Name, fld.Type.String())
}

func NewAnyType() *Type {
	return &Type{
		Kind:    TY_ANY,
		Name:    "any",
		MdlPath: "builtin",
	}
}

func NewVoidType() *Type {
	return &Type{
		Kind:    TY_VOID,
		Name:    "void",
		MdlPath: "builtin",
	}
}

func NewIntType() *Type {
	return &Type{
		Kind:    TY_INT,
		Name:    "int",
		MdlPath: "builtin",
	}
}

func NewFloatType() *Type {
	return &Type{
		Kind:    TY_FLOAT,
		Name:    "float",
		MdlPath: "builtin",
	}
}

func NewStringType() *Type {
	return &Type{
		Kind:    TY_STRING,
		Name:    "string",
		MdlPath: "builtin",
	}
}

func NewBoolType() *Type {
	return &Type{
		Kind:    TY_BOOL,
		Name:    "bool",
		MdlPath: "builtin",
	}
}

func NewNullType(child *Type) *Type {
	return &Type{
		Kind:    TY_NULL,
		Name:    "null",
		MdlPath: "builtin",
		Child:   child,
	}
}

func NewListType(child *Type) *Type {
	return &Type{
		Kind:    TY_LIST,
		Name:    "list",
		MdlPath: "builtin",
		Child:   child,
	}
}

func NewFuncType(modulePath string, args []*Type, ret *Type) *Type {
	return &Type{
		Kind:    TY_FUNC,
		Name:    "fn",
		MdlPath: modulePath,
		Args:    args,
		Return:  ret,
	}
}

func NewRecordType(name string, modulePath string, fields []FieldType) *Type {
	return &Type{
		Kind:    TY_RECORD,
		Name:    name,
		MdlPath: modulePath,
		Fields:  fields,
	}
}

func CompareTypes(a, b *Type) bool {
	if a == nil || b == nil {
		return false
	}
	if a.Kind == TY_ANY || b.Kind == TY_ANY {
		return true
	}
	if a.MdlPath != b.MdlPath {
		return false
	}
	switch a.Kind {
	case TY_VOID, TY_INT, TY_FLOAT, TY_BOOL, TY_STRING:
		return a.Kind == b.Kind
	case TY_NULL, TY_LIST:
		return a.Kind == b.Kind && CompareTypes(a.Child, b.Child)
	case TY_FUNC:
		if a.Kind != b.Kind || !CompareTypes(a.Return, b.Return) || len(a.Args) != len(b.Args) {
			return false
		}
		for i, argA := range a.Args {
			argB := b.Args[i]
			if !CompareTypes(argA, argB) {
				return false
			}
		}
		return true
	case TY_RECORD:
		if a.Name != b.Name || a.Kind != b.Kind || len(a.Fields) != len(b.Fields) {
			return false
		}
		for i, fldA := range a.Fields {
			fldB := b.Fields[i]
			if !CompareTypes(fldA.Type, fldB.Type) {
				return false
			}
		}
		return true
	}
	return false
}
