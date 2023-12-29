package dtype

import (
	"fmt"
	"strings"
)

type Type interface {
	Kind() TypeKind
	Name() string
	Compare(Type) bool
}

type TypeKind uint

const (
	Any TypeKind = 1 << iota
	Int
	Float
	Bool
	String
	Variatic
	List
	Function
)

func compare(t1, t2 Type) bool {
	return t1.Kind() == t2.Kind() && t1.Name() == t2.Name()
}

// =============================
type AnyType struct{}

func (*AnyType) Kind() TypeKind { return Any }
func (*AnyType) Name() string   { return "any" }
func (t *AnyType) Compare(other Type) bool {
	return true
}

type IntType struct{}

func (*IntType) Kind() TypeKind { return Int }
func (*IntType) Name() string   { return "int" }
func (t *IntType) Compare(other Type) bool {
	return compare(t, other)
}

type FloatType struct{}

func (*FloatType) Kind() TypeKind { return Float }
func (*FloatType) Name() string   { return "float" }
func (tp *FloatType) Compare(other Type) bool {
	return compare(tp, other)
}

type BoolType struct{}

func (*BoolType) Kind() TypeKind { return Bool }
func (*BoolType) Name() string   { return "bool" }
func (tp *BoolType) Compare(other Type) bool {
	return compare(tp, other)
}

type StringType struct{}

func (*StringType) Kind() TypeKind { return String }
func (*StringType) Name() string   { return "string" }
func (tp *StringType) Compare(other Type) bool {
	return compare(tp, other)
}

type VariaticType struct {
	Child Type
}

func (*VariaticType) Kind() TypeKind  { return Variatic }
func (tp *VariaticType) Name() string { return fmt.Sprintf("...%s", tp.Child.Name()) }
func (tp *VariaticType) Compare(other Type) bool {
	return compare(tp, other)
}

type ListType struct {
	Child Type
}

func (*ListType) Kind() TypeKind  { return List }
func (tp *ListType) Name() string { return fmt.Sprintf("list<%s>", tp.Child.Name()) }
func (tp *ListType) Compare(other Type) bool {
	return compare(tp, other)
}

type FunctionType struct {
	Args   []Type
	Return Type
}

func (*FunctionType) Kind() TypeKind { return Function }
func (t *FunctionType) Name() string {
	args := []string{}
	for _, arg := range t.Args {
		args = append(args, arg.Name())
	}
	ret := ""
	if t.Return != nil {
		ret = t.Return.Name()
	}
	return fmt.Sprintf("fn[%s]<%s>", strings.Join(args, " "), ret)
}
func (t *FunctionType) Compare(other Type) bool {
	return compare(t, other)
}
