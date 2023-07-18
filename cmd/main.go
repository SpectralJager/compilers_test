package main

import (
	"encoding/json"
	"fmt"
	"gl/internal"
	"os"

	"github.com/alecthomas/participle/v2"
)

func main() {
	code := `
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
}	
	`
	lex, err := internal.Def.LexString("", code)
	if err != nil {
		panic(err)
	}
	for {
		tok, err := lex.Next()
		if err != nil {
			panic(err)
		}
		if tok.EOF() {
			break
		}
		fmt.Println(tok.GoString())
	}
	fmt.Println()
	res, err := internal.Parser.ParseString("", code, participle.Trace(os.Stdout))
	fmt.Println()
	fmt.Println(internal.Parser.String())
	fmt.Println()
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
