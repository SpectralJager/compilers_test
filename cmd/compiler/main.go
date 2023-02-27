package main

import (
	"bytes"
	"encoding/json"
	"grimlang/internal/frontend"
	"grimlang/internal/frontend/ir"
	"log"

	"github.com/alecthomas/participle/v2"
)

func main() {
	code := `
	package main;

	fn main:void()
		while (lt i (len str))
			(println (string i))	
			(inc i)	
		end;
	end;
	`

	var buf bytes.Buffer
	res, err := frontend.Parser.ParseString("",
		code,
		participle.Trace(&buf),
	)
	if err != nil {
		log.Fatalf("%s\n%s", err, buf.String())
	}
	bts, _ := json.MarshalIndent(res, "", " ")
	log.Printf("#Parse Tree:\n%s", string(bts))
	// log.Printf("#BNF:\n%s", frontend.Parser.String())
	pChunk := ir.NewPackageChunk(res)
	log.Println(pChunk.Meta())
}
