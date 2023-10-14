package tp

import "fmt"

type Type struct {
	Kind    TypeKind
	SubType *Type
}

func NewInt() Type {
	return Type{
		Kind: Int,
	}
}

func NewBool() Type {
	return Type{
		Kind: Boolean,
	}
}

func NewVoid() Type {
	return Type{
		Kind: Void,
	}
}

func NewVariatic(subType *Type) Type {
	return Type{
		Kind:    Variatic,
		SubType: subType,
	}
}

func (t Type) String() string {
	switch t.Kind {
	case Void, Int, Boolean:
		return string(t.Kind)
	case Variatic:
		return fmt.Sprintf("...%s", t.SubType.String())
	default:
		return "undefined"
	}
}

func (t Type) Compare(other Type) bool {
	switch t.Kind {
	case Void, Int, Boolean:
		return t.Kind == other.Kind
	case Variatic:
		if t.Kind == other.Kind {
			return t.SubType.Compare(*other.SubType)
		}
		return false
	default:
		return false
	}
}

type TypeKind string

const (
	Void     TypeKind = "void"
	Int      TypeKind = "int"
	Boolean  TypeKind = "bool"
	Variatic TypeKind = "variatic"
)
