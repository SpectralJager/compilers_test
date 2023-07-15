package main

import (
	"encoding/json"
	"fmt"
	"gl/internal"
	"os"

	"github.com/alecthomas/participle/v2"
)

func main() {
	code := `
@import "github.com/alecthomas/participle" as participle
@import {
	"std" as std
	"fmt" as fmt
}

@const a:Int = 1	
@const {
	b:String = "hello"
	c:Float = 12.2
}

@var a:Int = 1	
@var {
	b:String = "hello"
	c:Float = 12.2
	d:Float = (fadd 12.2 13.0)
	list:List<Int> = '(12 13 14)
	map:Map<String,Int> = #("1"::1 "2"::2)
	complex:List<Map<String,Int>> = '(
		#("1"::1 "2"::2)
		#("3"::3 "4"::4))
}

@enum test:Int {
	first -> 1
	second -> 2
}

@enum test {
	first
	second
}
	`
	lex, err := internal.Def.LexString("", code)
	if err != nil {
		panic(err)
	}
	for {
		tok, err := lex.Next()
		if err != nil {
			panic(err)
		}
		if tok.EOF() {
			break
		}
		fmt.Println(tok.GoString())
	}
	fmt.Println()
	res, err := internal.Parser.ParseString("", code, participle.Trace(os.Stdout))
	fmt.Println()
	fmt.Println(internal.Parser.String())
	if err != nil {
		panic(err)
	}
	fmt.Println()
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
