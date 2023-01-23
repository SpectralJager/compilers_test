package main

import (
	"bytes"
	"fmt"
	"go/format"
	"grimlang/internal/core/backend/v3/compiler"
	"grimlang/internal/core/frontend/lexer"
	"grimlang/internal/core/frontend/parser"
)

func main() {
	code := `
	(fn bc [t a] (
		(def l 12)
		(def g 12)
		(def res (add l g t a))
		(ret res)
	))

	(fn main[] (
		(printf (bc 12 3))
	))
	`

	prog := parser.NewParser(lexer.NewLexer(code).Run()).Run()

	res := bytes.NewBufferString("")
	res.WriteString("package main\n\n")
	for _, v := range prog.Body {
		compiler.Compile(v, res)
	}

	// fmt.Println(res.String())
	bts, err := format.Source(res.Bytes())
	if err != nil {
		panic(err)
	}
	res.Reset()
	res.Write(bts)

	fmt.Println(res.String())
}
