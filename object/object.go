package object

import "grimlang/dtype"

type Object interface {
	Kind() ObjectKind
	Type() dtype.Type
}

type ObjectKind uint

const (
	Int ObjectKind = 3 << iota
	Float
	Bool
	String
	List
)

// =============================

type IntObject struct {
	Value int
}

func (*IntObject) Kind() ObjectKind { return Int }
func (*IntObject) Type() dtype.Type { return &dtype.IntType{} }

type FloatObject struct {
	Value float64
}

func (*FloatObject) Kind() ObjectKind { return Float }
func (*FloatObject) Type() dtype.Type { return &dtype.FloatType{} }

type BoolObject struct {
	Value bool
}

func (*BoolObject) Kind() ObjectKind { return Bool }
func (*BoolObject) Type() dtype.Type { return &dtype.BoolType{} }

type StringObject struct {
	Value string
}

func (*StringObject) Kind() ObjectKind { return String }
func (*StringObject) Type() dtype.Type { return &dtype.StringType{} }

type ListObject struct {
	ChildType dtype.Type
	Items     []Object
}

func (*ListObject) Kind() ObjectKind     { return List }
func (obj *ListObject) Type() dtype.Type { return &dtype.ListType{Child: obj.ChildType} }
