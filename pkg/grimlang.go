package main

import (
	"grimlang/pkg/repl"
	"os"
)

func main() {
	rpl := repl.Repl{}
	if len(os.Args) == 1 {
		repl.PrintHelp()
	} else if len(os.Args) == 2 {
		if os.Args[1] == "run" {
			rpl.RunRepl()
		} else {
			repl.PrintHelp()
		}
	} else if len(os.Args) == 3 {
		if os.Args[1] == "build" {
			rpl.Compile(os.Args[2])
		} else {
			repl.PrintHelp()
		}
	}
}
