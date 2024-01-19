package runtime

type Kind uint

const (
	Undefined Kind = iota

	Variable // variable symbol
	Constant // constant symbol
	Module   // module symbol
	Function // function symbol, type and litteral
	Record   // record symbol, type and litteral
	Void     // void type
	Any      // any type
	Null     // null type and litteral
	Int      // int type and litteral
	Float    // float type and litteral
	String   // string type and litteral
	Bool     // bool type and litteral
	List     // list type and litteral
	Builtin  // builtin function litteral
)

var kindNames = map[Kind]string{
	Undefined: "undefined",
	Variable:  "var",
	Constant:  "const",
	Module:    "module",
	Function:  "function",
	Record:    "record",
	Void:      "void",
	Any:       "any",
	Null:      "null",
	Int:       "int",
	Float:     "float",
	String:    "string",
	Bool:      "bool",
	List:      "list",
	Builtin:   "builtin",
}

func (kind Kind) String() string {
	name, ok := kindNames[kind]
	if !ok {
		return kindNames[Undefined]
	}
	return name
}

type Type interface {
	// general methods for all types
	Kind() Kind        // get kind of type
	Name() string      // get name of type
	Compare(Type) bool // compare types
	String() string
	// methods for specific types
	Subtype() Type // get list's items or null item type
	// function methods
	NumIn() int  // get number of ins
	In(int) Type // get function's i'th argument type
	Out() Type   // get function's return type
	// record methods
	NumField() int   // get number of record's fields
	Field(int) Field // get record's field by index
	FieldIndex(string) int
	FieldByName(string) Field // get record's field by name
}

type Field interface {
	Name() string
	Type() Type
}

type Litteral interface {
	Kind() Kind
	String() string
	Type() Type

	Int() int64
	Float() float64
	Bool() bool
	Str() string

	Item(int) Litteral
	Len() int

	Call(Context, ...Litteral) (Litteral, error)

	Field(int) Litteral
	FieldByName(string) Litteral
}

type Symbol interface {
	Kind() Kind
	Name() string
	String() string

	Type() Type
	Value() Litteral
	Set(Litteral) error
}

type Context interface {
	String() string
	Scope() string
	Search(string) Symbol
	Insert(Symbol) error
	Update(Symbol) error
}
