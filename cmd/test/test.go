package main

import (
	"encoding/json"
	"fmt"
	"gl/src/syntax"
)

func main() {
	source := `
@const test : int = 12;

@fn main() {}
`
	lexer := syntax.InitLexer(source)
	tokens := lexer.Run()
	if lexer.Errors() != nil {
		fmt.Println("----------Lexer Errors----------")
		for _, err := range lexer.Errors() {
			fmt.Println(err)
		}
		fmt.Println("--------------------------------")
	}
	parser := syntax.NewParser(tokens)
	program := parser.Run()
	if parser.Errors() != nil {
		fmt.Println("----------Parser Errors----------")
		for _, err := range parser.Errors() {
			fmt.Println(err)
		}
		fmt.Println("---------------------------------")
	}
	program.Filename = "test"
	res, _ := json.MarshalIndent(program, "", "   ")
	fmt.Printf("%s\n\n", res)
}
