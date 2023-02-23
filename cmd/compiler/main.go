package main

import (
	"bytes"
	"encoding/json"
	"grimlang/internal/frontend"
	"log"

	"github.com/alecthomas/participle/v2"
)

func main() {
	code := `
	package main;

	var
		x:int 10;
		y:int 12;
	end;

	fn main:void()
		"doc string"
		let i:int 10;
		let j:float 10.1;
		let str:string "some string";
		let temp:int (add i (sub 20 12));
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
	pChunk := frontend.NewPackageChunk(res)
	log.Println(pChunk.Meta())
}
