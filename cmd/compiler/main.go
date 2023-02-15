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
	// str, _ := json.MarshalIndent(res, "", " ")
	// fmt.Println(string(str))
	var buf bytes.Buffer
	err = backend.GenerateCode(&buf, *res)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(buf.String())
}
