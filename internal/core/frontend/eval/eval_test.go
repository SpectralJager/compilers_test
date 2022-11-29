package eval

import (
	"grimlang/internal/core/frontend/lexer"
	"grimlang/internal/core/frontend/parser"
	"testing"
)

func TestEval(t *testing.T) {
	code := `
	(add 1 2.2)
	(add 3 2)
	(add 3 (add 2 3))
	(add (add 12 8) (add 55 45 100) 200)
	(sub 10 5)
	(mul 10 10)
	(div 15 5)
	(def a 12)
	(a)
	(add a 18)
	(def pi 3.14)
	(pi)
	`

	tests := []struct {
		result interface{}
	}{
		{3},
		{5},
		{8},
		{420},
		{5},
		{100},
		{3},
		{"a"},
		{12},
		{30},
		{"pi"},
		{3.14},
	}
	env := make(Env)
	lex := lexer.NewLexer(code)
	pars := parser.NewParser(lex.Run())
	prog := pars.Run()
	for i, expRes := range tests {
		res, err := Eval(prog.Expresions[i], &env)
		if err != nil {
			t.Fatal(err)
		}
		if res != expRes.result {
			t.Fatalf("[%d] want: %v, got %v", i, expRes, res)
		}
	}
}
