package ir

/* Programm IR structure
.Filename: test
.Constants
	alpha int 12
	beta int 12
.Functions
; a:int, b:int -> int
_main:
	hlt
*/

type Program struct {
	Filename  string
	Constants []Constant
	Functions []Function
}

type Constant struct {
	Symbol string
	Value  Object
}

type Function struct {
	Symbol string
	Args   []Argument
	ReturnType
}
