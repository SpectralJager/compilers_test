package object

import (
	"fmt"
	"strings"
)

type DTypeAny struct{}
type DTypeInt struct{}
type DTypeFloat struct{}
type DTypeString struct{}
type DTypeBool struct{}

type DTypeList struct {
	ChildType Object
}

type DTypeNull struct {
	ChildType Object
}

type DTypeRecord struct {
	Identifier string
	FieldsType []Object
}

type DTypeFunction struct {
	ArgumentsType []Object
	ReturnType    Object
}

func (*DTypeAny) Kind() ObjectKind      { return AnyType }
func (*DTypeInt) Kind() ObjectKind      { return IntType }
func (*DTypeFloat) Kind() ObjectKind    { return FloatType }
func (*DTypeString) Kind() ObjectKind   { return StringType }
func (*DTypeBool) Kind() ObjectKind     { return BoolType }
func (*DTypeList) Kind() ObjectKind     { return ListType }
func (*DTypeNull) Kind() ObjectKind     { return NullType }
func (*DTypeRecord) Kind() ObjectKind   { return RecordType }
func (*DTypeFunction) Kind() ObjectKind { return FunctionType }

func (*DTypeAny) Inspect() string     { return "any" }
func (*DTypeInt) Inspect() string     { return "int" }
func (*DTypeFloat) Inspect() string   { return "float" }
func (*DTypeString) Inspect() string  { return "string" }
func (*DTypeBool) Inspect() string    { return "bool" }
func (dt *DTypeList) Inspect() string { return "list<" + dt.ChildType.Inspect() + ">" }
func (dt *DTypeNull) Inspect() string { return "null<" + dt.ChildType.Inspect() + ">" }
func (dt *DTypeRecord) Inspect() string {
	fields := []string{}
	for _, field := range dt.FieldsType {
		fields = append(fields, field.Inspect())
	}
	return fmt.Sprintf("%s{%s}", dt.Identifier, strings.Join(fields, " "))
}
func (dt *DTypeFunction) Inspect() string {
	args := []string{}
	for _, arg := range dt.ArgumentsType {
		args = append(args, arg.Inspect())
	}
	return fmt.Sprintf("fn[%s]<%s>", strings.Join(args, " "), dt.ReturnType.Inspect())
}
