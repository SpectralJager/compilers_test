package repl

import (
	"bufio"
	"fmt"
	"grimlang/internal/core/frontend/eval"
	"grimlang/internal/core/frontend/lexer"
	"grimlang/internal/core/frontend/parser"
	"os"
)

type Repl struct {
	replEnv eval.Env
}

func PrintHelp() {}

func (repl *Repl) RunRepl() {
	repl.replEnv = make(eval.Env)
	prompt := bufio.NewReader(os.Stdout)
	for {
		fmt.Print("repl-> ")
		req, err := prompt.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		repl.run(req)
		fmt.Print("\n")
	}
}

func (repl *Repl) Compile(src string) {}

func (repl *Repl) run(req string) {
	lex := lexer.NewLexer(req)
	toks := lex.Run()
	prs := parser.NewParser(toks)
	tree := prs.Run()
	result, err := eval.Eval(tree.Body[0], &repl.replEnv)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", result.ObjectValue())
}
