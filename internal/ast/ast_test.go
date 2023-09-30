package ast

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
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
	var errBuf strings.Builder
	programm, err := Parser.ParseString("test.grim", string(data), participle.Trace(&errBuf))
	if err != nil {
		fmt.Println(errBuf.String())
		t.Fatal(err)
	}
	res, _ := xml.MarshalIndent(programm, "", "  ")
	fmt.Println(string(res))
	fmt.Println(Parser.String())
}
