package main

import (
	"bytes"
	"fmt"
	"grimlang/internal/backend"
	"grimlang/internal/frontend"
	"log"
)

func main() {
	code := `
	(fn main:int []
		 "help doc"
		 (let a:int 12)
	)
	`
	res, err := frontend.Parser.ParseString("",
		code,
		// participle.Trace(os.Stdout),
	)
	if err != nil {
		log.Fatalf("%s", err)
	}
	var buf bytes.Buffer
	backend.GenerateProgram(&buf, res)
	fmt.Println(buf.String())
}
