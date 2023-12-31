package object

import (
	"fmt"
	"strings"
)

type DType interface {
	Object
	Compare(DType) bool
}

func compare(a, b DType) bool {
	return Is(a.Kind(), b.Kind()) && a.Inspect() == b.Inspect()
}

type DTypeAny struct{}
type DTypeInt struct{}
type DTypeFloat struct{}
type DTypeString struct{}
type DTypeBool struct{}
type DTypeList struct {
	ChildType DType
}
type DTypeNull struct {
	ChildType DType
}
type DTypeVariatic struct {
	ChildType DType
}
type DTypeRecord struct {
	FieldsType []DType
}
type DTypeFunction struct {
	ArgumentsType []DType
	ReturnType    DType
}

func (*DTypeAny) Kind() ObjectKind      { return AnyType }
func (*DTypeInt) Kind() ObjectKind      { return IntType }
func (*DTypeFloat) Kind() ObjectKind    { return FloatType }
func (*DTypeString) Kind() ObjectKind   { return StringType }
func (*DTypeBool) Kind() ObjectKind     { return BoolType }
func (*DTypeList) Kind() ObjectKind     { return ListType }
func (*DTypeNull) Kind() ObjectKind     { return NullType }
func (*DTypeVariatic) Kind() ObjectKind { return VariaticType }
func (*DTypeRecord) Kind() ObjectKind   { return RecordType }
func (*DTypeFunction) Kind() ObjectKind { return FunctionType }

func (*DTypeAny) Inspect() string         { return "any" }
func (*DTypeInt) Inspect() string         { return "int" }
func (*DTypeFloat) Inspect() string       { return "float" }
func (*DTypeString) Inspect() string      { return "string" }
func (*DTypeBool) Inspect() string        { return "bool" }
func (dt *DTypeList) Inspect() string     { return "list<" + dt.ChildType.Inspect() + ">" }
func (dt *DTypeNull) Inspect() string     { return "null<" + dt.ChildType.Inspect() + ">" }
func (dt *DTypeVariatic) Inspect() string { return "..." + dt.ChildType.Inspect() }
func (dt *DTypeRecord) Inspect() string {
	fields := []string{}
	for _, field := range dt.FieldsType {
		fields = append(fields, field.Inspect())
	}
	return fmt.Sprintf("record{%s}", strings.Join(fields, " "))
}
func (dt *DTypeFunction) Inspect() string {
	args := []string{}
	for _, arg := range dt.ArgumentsType {
		args = append(args, arg.Inspect())
	}
	return fmt.Sprintf("fn[%s]<%s>", strings.Join(args, " "), dt.ReturnType.Inspect())
}

func (tp *DTypeAny) Compare(other DType) bool {
	return compare(tp, other)
}
func (tp *DTypeInt) Compare(other DType) bool {
	return compare(tp, other)
}
func (tp *DTypeFloat) Compare(other DType) bool {
	return compare(tp, other)
}
func (tp *DTypeString) Compare(other DType) bool {
	return compare(tp, other)
}
func (tp *DTypeBool) Compare(other DType) bool {
	return compare(tp, other)
}
func (tp *DTypeList) Compare(other DType) bool {
	return compare(tp, other)
}
func (tp *DTypeNull) Compare(other DType) bool {
	return compare(tp, other)
}
func (tp *DTypeVariatic) Compare(other DType) bool {
	return compare(tp, other)
}
func (tp *DTypeRecord) Compare(other DType) bool {
	return compare(tp, other)
}
func (tp *DTypeFunction) Compare(other DType) bool {
	return compare(tp, other)
}
