package runtime

/*
-----------<Enviroment
builtin |>

	int -> import(int)
	main -> import(main)
	exit -> fn[int] <void>
	...

int |>

	add -> fn[...int] <int>
	sub -> fn[...int] <int>
	...

main::builtin|>

	vector -> import(hash1)
	main -> fn[] <void>

hash1::builtin|>

	point -> import(hash2)
	vector -> record<p0::point/point p1::point/point>
	new_vector -> fn[point/point point/point] <vector>
	...

hash2::builtin|>

	point -> record<x0::int x1::int>
	new_point -> fn[int int] <point>
	...
*/
type Kind uint

const (
	Invalid Kind = iota

	ImportSymbol
	VariableSymbol
	ConstantSymbol
	FunctionSymbol
	RecordSymbol

	IntLitteral
	RecordLitteral
	FunctionLitteral
	BuiltinLitteral

	VoidType
	IntType
	RecordType
	FunctionType
)

type Enviroment interface {
	String() string
	Tag() string
	Parent() Enviroment
	Search(string) Symbol
	SearchLocal(string) Symbol

	Insert(Symbol) error
}

type Symbol interface {
	String() string
	Kind() Kind
	Name() string
	Type() Type
	Value() Litteral
	Tag() string

	Set(Litteral) error
}

type Litteral interface {
	String() string
	Kind() Kind
	Type() Type
	Int() int64
	Field(int) Litteral
	FieldByName(string) Litteral
}

type Type interface {
	String() string
	Kind() Kind
	Name() string
	NumIns() int
	In(int) Type
	Out() Type
	Field(int) Field
	FieldByName(string) Field
	FieldIndex(string) int
	NumFields() int
}

type Field interface {
	String() string
	Name() string
	Type() Type
}

type enviroment struct {
	parent  Enviroment
	tag     string
	symbols map[string]Symbol
}
