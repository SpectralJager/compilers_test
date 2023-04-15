package main

import (
	"fmt"
	"gl/src/syntax"
)

func main() {
	source := `
some test
token
12
12.1
`
	lexer := syntax.InitLexer(source)
	tokens := lexer.Run()
	if lexer.Errors() != nil {
		fmt.Println("----------Errors----------")
		for _, err := range lexer.Errors() {
			fmt.Println(err)
		}
		fmt.Println("--------------------------")
	}
	for _, token := range tokens {
		fmt.Println(token)
	}
}
