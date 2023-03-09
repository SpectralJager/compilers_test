package main

import (
	"bytes"
	"encoding/json"
	"grimlang/internal/frontend"
	"grimlang/internal/ir"
	"log"

	"github.com/alecthomas/participle/v2"
)

func main() {
	code := `
	var 
		i:int 10;
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
	res.Filename = "TestFile"
	bts, _ := json.MarshalIndent(res, "", " ")
	log.Printf("#Parse Tree:\n%s", string(bts))
	// log.Printf("#BNF:\n%s", frontend.Parser.String())
	cCode := ir.NewPackageChunk(res)
	log.Println(cCode.Meta())
}
