package ir

import (
	"fmt"
	"testing"
)

func TestIR(t *testing.T) {
	ir := Program{
		Name: "test",
		Constants: []IConstant{
			&Integer{40},
			&Integer{2},
			&Integer{1},
		},
		Globals: map[string]ISymbolDef{
			"fib": &FunctionDef{
				Name: "fib",
				Arguments: []IDataType{
					&Primitive{"int"},
				},
				Returns: []IDataType{
					&Primitive{"int"},
				},
			},
			"main": &FunctionDef{
				Name: "main",
				Arguments: []IDataType{
					&Primitive{"void"},
				},
				Returns: []IDataType{
					&Primitive{"void"},
				},
			},
		},
		Functions: map[string]Function{
			"main": {
				Name:      "main",
				Arguments: map[string]ISymbolDef{},
				Locals:    map[string]ISymbolDef{},
				BodyCode: []IInstruction{
					&Load{ConstIndex: 0},
					&Call{FuncName: "fib"},
					&CallBuildin{FuncName: "print"},
					&Return{},
				},
			},
			"fib": {
				Name: "fib",
				Arguments: map[string]ISymbolDef{
					"n": &VaribleDef{
						Name: "n",
						Type: &Primitive{"int"},
					},
				},
				Locals: map[string]ISymbolDef{},
				BodyCode: []IInstruction{
					&Load{0},
					&LocalLoad{"n"},
					&CallBuildin{"lt"},
					&ConditionalJump{14},
					&Load{1},
					&LocalLoad{"n"},
					&CallBuildin{"sub"},
					&Call{"fib"},
					&Load{2},
					&LocalLoad{"n"},
					&CallBuildin{"sub"},
					&Call{"fib"},
					&CallBuildin{"add"},
					&Return{1},
					&LocalLoad{"n"},
					&Return{1},
				},
			},
		},
	}
	fmt.Println(ir.String())
}
