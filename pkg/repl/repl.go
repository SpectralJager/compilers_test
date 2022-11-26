package repl

import (
	"bufio"
	"fmt"
	"grimlang/internal/core/frontend/lexer"
	"grimlang/internal/core/frontend/parser"
	"os"
)

type Repl struct {
}

func PrintHelp() {}

func RunRepl() {
	prompt := bufio.NewReader(os.Stdout)
	for {
		fmt.Print("repl-> ")
		req, err := prompt.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		run(req)
		fmt.Print("\n")
	}
}

func Compile(src string) {}

func run(req string) {
	lex := lexer.NewLexer(req)
	toks := lex.Run()
	prs := parser.NewParser(toks)
	tree := prs.Run()
	for _, expr := range tree.Expresions {
		fmt.Printf("%+v\n", expr)
	}
}
