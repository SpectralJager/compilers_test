package repl

import (
	"bufio"
	"fmt"
	"grimlang/internal/core/frontend/lexer"
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
	for _, tok := range toks {
		fmt.Printf("%v\n", tok.String())
	}
}
