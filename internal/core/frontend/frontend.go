package main

import (
	"fmt"
	"grimlang/internal/core/frontend/lexer"
	"grimlang/internal/core/frontend/parser"
)

func main() {
	code := `
	(fn recur [c] (
		(if (leq c 0) (
			(ret)
		))
		(println c)
		(recur (sub c 1))

	))
	(def cnt 10)
	(recur cnt)
	`
	programm := parser.NewParser(
		lexer.NewLexer(code).Run(),
	).Run()
	fmt.Println(programm)
}
