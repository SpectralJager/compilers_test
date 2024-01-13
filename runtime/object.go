package runtime

import (
	"fmt"
	"grimlang/ast"
	"strings"
)

type ObjectKind uint

const (
	OB_NIL ObjectKind = iota
	OB_INT
	OB_FLOAT
	OB_BOOL
	OB_STRING
	OB_LIST
	OB_FUNC
	OB_BUILTIN
	OB_RECORD
)

type Object struct {
	// general object fields
	Kind ObjectKind
	Type *Type

	// primitive objects
	I int64
	F float64
	B bool
	S string

	// list object
	Items []*Object

	// user defined function
	Fn *ast.FunctionDecl
	// builtin function
	Builtin func(args ...*Object) (*Object, error)

	// record object
	Fields []FieldObject
}

type FieldObject struct {
	Name  string
	Type  *Type
	Value *Object
}

func (obj *Object) String() string {
	switch obj.Kind {
	case OB_NIL:
		return "nil"
	case OB_INT:
		return fmt.Sprintf("%d", obj.I)
	case OB_FLOAT:
		return fmt.Sprintf("%f", obj.F)
	case OB_BOOL:
		return fmt.Sprintf("%v", obj.B)
	case OB_STRING:
		return obj.S
	case OB_LIST:
		items := []string{}
		for _, object := range obj.Items {
			items = append(items, object.String())
		}
		return fmt.Sprintf("list<%s>{%s}", obj.Type, strings.Join(items, " "))
	case OB_FUNC, OB_BUILTIN:
		return fmt.Sprintf("%s", obj.Type)
	case OB_RECORD:
		fields := []string{}
		for _, fld := range obj.Fields {
			fields = append(fields, fld.String())
		}
		return fmt.Sprintf("record{%s}", strings.Join(fields, " "))
	}
	return ""
}

func (fldObj *FieldObject) String() string {
	return fmt.Sprintf("%s::(%s)%s", fldObj.Name, fldObj.Type, fldObj.Value)
}

func NewNilObject() *Object {
	return &Object{
		Kind: OB_NIL,
		Type: NewAnyType(),
	}
}

func NewIntObject(val int64) *Object {
	return &Object{
		Kind: OB_INT,
		I:    val,
		Type: NewIntType(),
	}
}

func NewFloatObject(val float64) *Object {
	return &Object{
		Kind: OB_FLOAT,
		F:    val,
		Type: NewFloatType(),
	}
}

func NewBoolObject(val bool) *Object {
	return &Object{
		Kind: OB_BOOL,
		B:    val,
		Type: NewBoolType(),
	}
}

func NewStringObject(val string) *Object {
	return &Object{
		Kind: OB_STRING,
		S:    val,
		Type: NewStringType(),
	}
}

func NewListObject(childType *Type, items ...*Object) *Object {
	for _, item := range items {
		if !CompareTypes(childType, item.Type) {
			panic(fmt.Errorf("new list: can't use item of %s in list of %s", item.Type.String(), childType.String()))
		}
	}
	return &Object{
		Kind:  OB_LIST,
		Items: items,
		Type:  NewListType(childType),
	}
}

func NewFuncObject(typ *Type, fn *ast.FunctionDecl) *Object {
	return &Object{
		Kind: OB_FUNC,
		Fn:   fn,
		Type: typ,
	}
}

func NewBuiltinObject(typ *Type, builtin func(args ...*Object) (*Object, error)) *Object {
	return &Object{
		Kind:    OB_BUILTIN,
		Builtin: builtin,
		Type:    typ,
	}
}
