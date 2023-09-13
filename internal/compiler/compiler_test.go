package compiler

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/participle/v2"
)

func TestLexer(t *testing.T) {
	data, err := os.ReadFile("../../test.grim")
	if err != nil {
		t.Fatal(err)
	}
	lex, err := Lexer.LexString("test.grim", string(data))
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
		fmt.Println(tok.GoString())
	}
}

func TestParser(t *testing.T) {
	data, err := os.ReadFile("../../test.grim")
	if err != nil {
		t.Fatal(err)
	}
	programm, err := Parser.ParseString("test.grim", string(data), participle.Trace(os.Stderr))
	if err != nil {
		t.Fatal(err)
	}
	res, _ := json.MarshalIndent(programm, "", "  ")
	fmt.Println(string(res))
}

func TestCodegen(t *testing.T) {
	data, err := os.ReadFile("../../test.grim")
	if err != nil {
		t.Fatal(err)
	}
	programm, err := Parser.ParseString("test.grim", string(data))
	if err != nil {
		t.Fatal(err)
	}
	res := Generate(programm)
	fmt.Println(res)
}
