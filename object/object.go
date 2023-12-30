package object

type ObjectKind uint

const (
	_ ObjectKind = 1 << iota
	VariableSymbol
	ConstantSymbol
	FunctionSymbol
	BuiltinSymbol
	ModuleSymbol
	RecordSymbol

	AnyType
	NullType
	IntType
	FloatType
	StringType
	BoolType
	ListType
	RecordType
	FunctionType

	NullLitteral
	IntLitteral
	FloatLitteral
	StringLitteral
	BoolLitteral
	ListLitteral
	RecordLitteral

	IsSymbol   = VariableSymbol | ConstantSymbol | FunctionSymbol | BuiltinSymbol | ModuleSymbol | RecordSymbol
	IsType     = AnyType | NullType | IntType | FloatType | StringType | BoolType | ListType | RecordType | FunctionType
	IsLitteral = NullLitteral | IntLitteral | FloatLitteral | StringLitteral | BoolLitteral | ListLitteral | RecordLitteral
)

func Is(a, b ObjectKind) bool {
	return a&b != 0
}

type Object interface {
	Kind() ObjectKind
	Inspect() string
}
