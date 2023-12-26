package object

import "grimlang/dtype"

type Object interface {
	Kind() ObjectKind
	Type() dtype.Type
}

type ObjectKind uint

const (
	Int ObjectKind = 3 << iota
)

// =============================

type IntObject struct {
	Value int
}

func (*IntObject) Kind() ObjectKind { return Int }
func (*IntObject) Type() dtype.Type { return &dtype.IntType{} }
