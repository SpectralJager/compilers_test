package analysis

import tp "grimlang/internal/type"

var Builtins = SymbolTable{
	&FunctionSymbol{
		Ident:       "int/add",
		Args:        []tp.Type{tp.VariaticType{Subtype: tp.IntegerType{}}},
		ReturnTypes: []tp.Type{tp.IntegerType{}},
	},
	&FunctionSymbol{
		Ident:       "int/sub",
		Args:        []tp.Type{tp.VariaticType{Subtype: tp.IntegerType{}}},
		ReturnTypes: []tp.Type{tp.IntegerType{}},
	},
	&FunctionSymbol{
		Ident:       "int/mul",
		Args:        []tp.Type{tp.VariaticType{Subtype: tp.IntegerType{}}},
		ReturnTypes: []tp.Type{tp.IntegerType{}},
	},
	&FunctionSymbol{
		Ident:       "int/div",
		Args:        []tp.Type{tp.VariaticType{Subtype: tp.IntegerType{}}},
		ReturnTypes: []tp.Type{tp.IntegerType{}},
	},
	&FunctionSymbol{
		Ident:       "int/lt",
		Args:        []tp.Type{tp.VariaticType{Subtype: tp.IntegerType{}}},
		ReturnTypes: []tp.Type{tp.BooleanType{}},
	},
	&FunctionSymbol{
		Ident:       "int/gt",
		Args:        []tp.Type{tp.VariaticType{Subtype: tp.IntegerType{}}},
		ReturnTypes: []tp.Type{tp.BooleanType{}},
	},
	&FunctionSymbol{
		Ident:       "int/eq",
		Args:        []tp.Type{tp.VariaticType{Subtype: tp.IntegerType{}}},
		ReturnTypes: []tp.Type{tp.BooleanType{}},
	},
	&FunctionSymbol{
		Ident:       "int/leq",
		Args:        []tp.Type{tp.VariaticType{Subtype: tp.IntegerType{}}},
		ReturnTypes: []tp.Type{tp.BooleanType{}},
	},
	&FunctionSymbol{
		Ident:       "int/geq",
		Args:        []tp.Type{tp.VariaticType{Subtype: tp.IntegerType{}}},
		ReturnTypes: []tp.Type{tp.BooleanType{}},
	},
}
