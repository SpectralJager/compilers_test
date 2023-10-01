package lexer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	testCases := []struct {
		src string
		exp []string
	}{
		{
			src: "@var a:int = 12;",
			exp: []string{"@var", "a", ":", "int", "=", "12", ";"},
		},
		{
			src: "@var alca12:int = -12;",
			exp: []string{"@var", "alca12", ":", "int", "=", "-12", ";"},
		},
		{
			src: "@var _alca_12:int = +12;",
			exp: []string{"@var", "_alca_12", ":", "int", "=", "+12", ";"},
		},
		{
			src: "@fn func() {}",
			exp: []string{"@fn", "func", "(", ")", "{", "}"},
		},
		{
			src: "@fn func(a:int b:int) {}",
			exp: []string{"@fn", "func", "(", "a", ":", "int", "b", ":", "int", ")", "{", "}"},
		},
		{
			src: "@fn _func_12(a:int b:int) {}",
			exp: []string{"@fn", "_func_12", "(", "a", ":", "int", "b", ":", "int", ")", "{", "}"},
		},
		{
			src: "@fn _func_12(a:int b:int) { @var a:int = 12; }",
			exp: []string{"@fn", "_func_12", "(", "a", ":", "int", "b", ":", "int", ")", "{", "@var", "a", ":", "int", "=", "12", ";", "}"},
		},
		{
			src: "(int/add a 12)",
			exp: []string{"(", "int", "/", "add", "a", "12", ")"},
		},
		{
			src: "@while (int/lt a 12) {}",
			exp: []string{"@while", "(", "int", "/", "lt", "a", "12", ")", "{", "}"},
		},
	}
	for i, tC := range testCases {
		t.Run(tC.src, func(t *testing.T) {
			l, err := Lexer.LexString(fmt.Sprint(i), tC.src)
			if err != nil {
				t.Fatalf("$%d: can't parse '%s': %s", i, tC.src, err.Error())
			}
			var exp []string
			for tok, err := l.Next(); !tok.EOF(); tok, err = l.Next() {
				if err != nil {
					t.Fatalf("$%d: can't get token: %s", i, err.Error())
				}
				exp = append(exp, tok.Value)
			}
			if !reflect.DeepEqual(tC.exp, exp) {
				t.Fatalf("$%d: expected %v, got %v", i, tC.exp, exp)
			}
		})
	}
}
