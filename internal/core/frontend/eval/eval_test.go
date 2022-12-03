package eval

import (
	"grimlang/internal/core/frontend/lexer"
	"grimlang/internal/core/frontend/parser"
	"testing"
)

func TestEval(t *testing.T) {
	code := `
	12
	`

	tests := []struct {
		result interface{}
	}{
		{12},
	}
	env := make(Env)
	lex := lexer.NewLexer(code)
	pars := parser.NewParser(lex.Run())
	prog := pars.Run()
	for i, expRes := range tests {
		res, err := Eval(prog.Body[i], &env)
		if err != nil {
			t.Fatalf("[%d] %s", i, err)
		}
		if res != expRes.result {
			t.Fatalf("[%d] want: %v, got %v", i, expRes, res)
		}
	}
}
