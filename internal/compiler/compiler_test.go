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
	code := `
@var result:int = 0;

@fn fib[n:int] <int> {
    @if (ilt n 2) {
        @return n;
    }
    @return (iadd 
        (fib (isub n 1)) 
        (fib (isub n 2)));
}

@fn main[] {
   @set result = (fib 40);
}`
	// @var result:int = 0;
	// @fn main[] {
	// 	@var temp:int = 0;
	// 	@if (ieq temp 0) {
	// 		@set result = 10;
	// 	} else {
	// 		@set result = 20;
	// 	}
	// }
	// 	`
	programm, err := Parser.ParseString("test.grim", string(code))
	if err != nil {
		t.Fatal(err)
	}
	var g Generator
	res, err := g.GenerateProgram(*programm)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.String())
}
