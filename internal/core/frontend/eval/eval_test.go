package eval

import (
	"grimlang/internal/core/frontend/lexer"
	"grimlang/internal/core/frontend/parser"
	"testing"
)

func TestEval(t *testing.T) {
	code := `
	(add 1 2)
	(add 3 2)
	(add 3 (add 2 3))
	`

	tests := []struct {
		result interface{}
	}{
		{3},
		{5},
		{8},
	}
	env := new(Env)
	lex := lexer.NewLexer(code)
	pars := parser.NewParser(lex.Run())
	prog := pars.Run()
	for i, expRes := range tests {
		res, err := Eval(prog.Expresions[i], env)
		if err != nil {
			t.Fatal(err)
		}
		if res != expRes.result {
			t.Fatalf("[%d] want: %v, got %v", i, expRes, res)
		}
	}
}
