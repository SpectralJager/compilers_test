package parser

import (
	"gl/core/frontend/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	code := `
@const a:int = 1;
@var b:int = (add 2 a);
@fn main:void() {
	@const a:int = 1;
	@set b = a;
	@while (neq b 10) {
		@set b = (add b 1);
	}
	@if (neg b 10) {
		(printf "%d" b)
	} else {
		(printf "%d" (sub b 3))
	}
}
	`
	lex := lexer.NewLexer(code)
	tokens := lex.Lex()
	if len(lex.Error()) != 0 {
		for _, e := range lex.Error() {
			t.Logf("Error: %v", e)
		}
		t.FailNow()
	}
	prs := NewParser(*tokens)
	_ = prs.Parse()
	if len(prs.Errors()) != 0 {
		for _, e := range prs.Errors() {
			t.Logf("Error: %v", e)
		}
		t.FailNow()
	}
	// data, _ := json.MarshalIndent(programm, "", "  ")
	// t.Fatalf("%s\n", data)
}

func TestParseConst(t *testing.T) {
	td := []struct {
		code  string
		isNil bool
	}{
		{`@const alpha:int = 12;`, false},
		{`@const beta:float = 12.2;`, false},
		{`@const gamma:string = "gamma string";`, false},
		{`@const teta:bool = true;`, false},
		{`@const delta:bool = false;`, false},
		{`@const buba:int = (add 1 2);`, true},
		{`@const :int = 22;`, true},
		{`@const fiz = 12;`, true},
		{`@set fiz = 12;`, true},
	}

	for _, tc := range td {
		t.Run(tc.code, func(t *testing.T) {
			parser := NewParser(*lexer.NewLexer(tc.code).Lex())
			result := parser.parseConst()
			if (result == nil) != tc.isNil {
				t.Logf("Expected to be nil -- %v, got %v", tc.isNil, result)
				for _, err := range parser.Errors() {
					t.Logf("%s\n", err)
				}
				t.FailNow()
			}
		})
	}
}

func TestParseVar(t *testing.T) {
	td := []struct {
		code  string
		isNil bool
	}{
		{`@var alpha:int = 12;`, false},
		{`@var beta:float = 12.2;`, false},
		{`@var gamma:string = "gamma string";`, false},
		{`@var teta:bool = true;`, false},
		{`@var delta:bool = false;`, false},
		{`@var buba:int = (add 1 2);`, false},
		{`@var :int = 22;`, true},
		{`@var fiz = 12;`, true},
		{`@set fiz = 12;`, true},
	}

	for _, tc := range td {
		t.Run(tc.code, func(t *testing.T) {
			parser := NewParser(*lexer.NewLexer(tc.code).Lex())
			result := parser.parseVar()
			if (result == nil) != tc.isNil {
				t.Logf("Expected to be nil -- %v, got %v", tc.isNil, result)
				for _, err := range parser.Errors() {
					t.Logf("%s\n", err)
				}
				t.FailNow()
			}
		})
	}
}

func TestParseSet(t *testing.T) {
	td := []struct {
		code  string
		isNil bool
	}{
		{`@set alpha = 12;`, false},
		{`@set beta = 12.2;`, false},
		{`@set gamma = "gamma string";`, false},
		{`@set teta = true;`, false},
		{`@set delta = false;`, false},
		{`@set buba = (add 1 2);`, false},
		{`@set buba:int = 22;`, true},
		{`@set fiz = ;`, true},
		{`@const fiz = 12;`, true},
	}

	for _, tc := range td {
		t.Run(tc.code, func(t *testing.T) {
			parser := NewParser(*lexer.NewLexer(tc.code).Lex())
			result := parser.parseSet()
			if (result == nil) != tc.isNil {
				t.Logf("Expected to be nil -- %v, got %v", tc.isNil, result)
				for _, err := range parser.Errors() {
					t.Logf("%s\n", err)
				}
				t.FailNow()
			}
		})
	}
}

func TestParseIf(t *testing.T) {
	td := []struct {
		code  string
		isNil bool
	}{
		{`@if true {(printf "hi")}`, false},
		{`@if false {(printf "hi")} else {(printf str)}`, false},
		{`@if (gt 12 2) {(printf "hi")}`, false},
		{`@if (eq 12 2)  {(printf "hi")} else {(printf str)}`, false},
		{`@if (eq 12 12)  {@const alpha:int = 12;} else {(printf str)}`, false},
		{`@if (eq 12 12)  {@fn main:int(){}} else {(printf str)}`, true},
		{`@if {} else {(printf str)}`, true},
	}

	for _, tc := range td {
		t.Run(tc.code, func(t *testing.T) {
			parser := NewParser(*lexer.NewLexer(tc.code).Lex())
			result := parser.parseIf()
			if (result == nil) != tc.isNil {
				t.Logf("Expected to be nil -- %v, got %v", tc.isNil, result)
				for _, err := range parser.Errors() {
					t.Logf("%s\n", err)
				}
				t.FailNow()
			}
		})
	}
}

func TestParseWhile(t *testing.T) {
	td := []struct {
		code  string
		isNil bool
	}{
		{`@while true {(printf "hi")}`, false},
		{`@while false {(printf "hi")}`, false},
		{`@while (gt 12 2) {(printf "hi")}`, false},
		{`@while (eq 12 12)  {@const alpha:int = 12;}`, false},
		{`@while (eq 12 12)  {@fn main:int(){}}`, true},
		{`@while {} `, true},
	}

	for _, tc := range td {
		t.Run(tc.code, func(t *testing.T) {
			parser := NewParser(*lexer.NewLexer(tc.code).Lex())
			result := parser.parseWhile()
			if (result == nil) != tc.isNil {
				t.Logf("Expected to be nil -- %v, got %v", tc.isNil, result)
				for _, err := range parser.Errors() {
					t.Logf("%s\n", err)
				}
				t.FailNow()
			}
		})
	}
}

func TestParseFn(t *testing.T) {
	td := []struct {
		code  string
		isNil bool
	}{
		{`@fn main:void (){(printf "hi")}`, false},
		{`@fn main:void (a:int){(printf "hi")}`, false},
		{`@fn main:void (a:int b:float){(printf "hi")}`, false},
		{`@fn main:void (){@const alpha:int = 12;@set alpha = 22;}`, false},
		{`@fn main:void (){@fn main:int(){}}`, true},
		{`@fn main (){}`, true},
		{`@fn main:void (a b:int){}`, true},
		{`@fn :void (a:int b:int){}`, true},
		{`@const main:void (a:int b:int){}`, true},
	}

	for _, tc := range td {
		t.Run(tc.code, func(t *testing.T) {
			parser := NewParser(*lexer.NewLexer(tc.code).Lex())
			result := parser.parseFunction()
			if (result == nil) != tc.isNil {
				t.Logf("Expected to be nil -- %v, got %v", tc.isNil, result)
				for _, err := range parser.Errors() {
					t.Logf("%s\n", err)
				}
				t.FailNow()
			}
		})
	}
}

func TestParseExpression(t *testing.T) {
	td := []struct {
		code  string
		isNil bool
	}{
		{`(add 1 2)`, false},
		{`(add 1 (mul 2 2))`, false},
		{`(add 1 (mul 2 2) (div 3 2))`, false},
		{`(someFunction)`, false},
		{`add 1 2)`, true},
		{`add`, true},
		{`(add 1 @const a:int = 12;)`, true},
	}

	for _, tc := range td {
		t.Run(tc.code, func(t *testing.T) {
			parser := NewParser(*lexer.NewLexer(tc.code).Lex())
			result := parser.parseExpression()
			if (result == nil) != tc.isNil {
				t.Logf("Expected to be nil -- %v, got %v", tc.isNil, result)
				for _, err := range parser.Errors() {
					t.Logf("%s\n", err)
				}
				t.FailNow()
			}
		})
	}
}
