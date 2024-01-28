package runtime

/*
-----------<Environment
builtin |>
	int -> import(int)
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
	Undefined Kind = iota
	// SYmbols
	SY_Variable
	SY_Constant
	SY_Function
	SY_Builtin
	SY_Import
	SY_Record
	// TYpes
	TY_Any
	TY_Void
	TY_Int
	TY_Float
	TY_String
	TY_Bool
	TY_List
	TY_Variatic
	TY_Function
	TY_Record
	// LItterals
	LI_Int
	LI_Float
	LI_String
	LI_Bool
	LI_List
	LI_Record
)
