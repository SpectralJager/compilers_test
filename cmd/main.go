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

@struct Test {
	a:Int
	b:String
	complex:List<Map<String,Int>>
	
	@fn print:Void(a:Int b:String complex:List<Map<String,Int>>){
		(printf "a:%d b:%s" a b)
	}
}

@fn main:Int(a:Int b:String complex:List<Map<String,Int>>) {
	(printf "a:%d b:%s" a b)
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
	fmt.Println()
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
