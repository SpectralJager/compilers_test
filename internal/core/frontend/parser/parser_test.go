package parser

import (
	"grimlang/internal/core/frontend/ast"
	"grimlang/internal/core/frontend/lexer"
	"grimlang/internal/core/frontend/tokens"
	"strings"
	"testing"
)

func TestParseSExpression(t *testing.T) {
	tests := []struct {
		input                string
		expectedStartlistLit string
		expectedOperatorLit  string
		expectedEndlistLit   string
		expectedArgumentsLit []string
	}{
		{"(+ b c)", "(", "+", ")", []string{"b", "c"}},
		{"(add 1 12)", "(", "add", ")", []string{"1", "12"}},
		{"(sum 1 12)", "(", "sum", ")", []string{"1", "12"}},
		{"(test)", "(", "test", ")", []string{}},
	}

	for _, tt := range tests {
		ch := make(chan tokens.Token)
		lexer.NewLexer(strings.NewReader(tt.input), ch)
		pr := NewParser(ch)
		res := pr.ParseProgram().Nodes[0]
		if res.TokenLiteral() != tt.input {
			t.Errorf("res.TokenLiteral not '%q'. got=%q", tt.input, res.TokenLiteral())
		}

		se, ok := res.(*ast.SExpression)
		if !ok {
			t.Errorf("res not %T. got=%T", &ast.SExpression{}, se)
		}

		if se.StartListToken.Literal != tt.expectedStartlistLit {
			t.Errorf("se.StartListToken.Literal not '%s'. got=%s", tt.expectedStartlistLit, se.StartListToken.Literal)
		}
		if se.EndListToken.Literal != tt.expectedEndlistLit {
			t.Errorf("se.EndListToken.Literal not '%s'. got=%s", tt.expectedEndlistLit, se.EndListToken.Literal)
		}
		if se.Operation.Literal != tt.expectedOperatorLit {
			t.Errorf("se.Operation.Literal not '%s'. got=%s", tt.expectedOperatorLit, se.Operation.Literal)
		}
		for i, arg := range se.Arguments {
			if arg.TokenLiteral() != tt.expectedArgumentsLit[i] {
				t.Errorf("arg #%d not '%s'. got=%s", i, tt.expectedArgumentsLit[i], arg.TokenLiteral())
			}
		}

	}
}
