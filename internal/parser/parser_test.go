package parser

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/alecthomas/participle/v2"
)

var code = `
// this is comment
@var global:int = 1;
@fn main(a:int) <int> {
	@var local:int = 2;
	@var local_list:list<int> = '(1 2 3);
	@set local = 10;
	(function)
	@while (eq a b) {
		(iadd a b)
	}
	@each i:int <- local_list {
		(print i)
	}
	@if (eq a b) {
		(print i)
	}
}`

func TestLexer(t *testing.T) {
	lex, err := Lexer.LexString("", code)
	if err != nil {
		t.Fatal(err)
	}
	for {
		tok, err := lex.Next()
		if err != nil {
			t.Fatal(err)
		}
		if tok.EOF() {
			break
		}
		t.Logf("%s\n", tok.GoString())
	}
}

func TestParser(t *testing.T) {
	var buffer bytes.Buffer
	t.Logf("\n%s\n", Parser.String())
	res, err := Parser.ParseString("", code, participle.Trace(&buffer))
	if err != nil {
		t.Log(buffer.String())
		t.Fatal(err)
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
