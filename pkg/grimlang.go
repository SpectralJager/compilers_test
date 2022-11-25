package main

import (
	"grimlang/pkg/repl"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		repl.PrintHelp()
	} else if len(os.Args) == 2 {
		if os.Args[1] == "run" {
			repl.RunRepl()
		} else {
			repl.PrintHelp()
		}
	} else if len(os.Args) == 3 {
		if os.Args[1] == "build" {
			repl.Compile(os.Args[2])
		} else {
			repl.PrintHelp()
		}
	}
}
