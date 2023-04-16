package main

import (
	"fmt"
	"gl/src/syntax"
)

func main() {
	source := `
@import "std" as std;
@import "http" as http;

const a:i32 = 1;

pub fn main() void {
	@var server:http/Server = (http/NewServer);
	@set server/address = "localhost";
	@set server/port = 8080;

	@var index:http/Handler = (http/NewHandler "/" );

	(server/serve)
}

`
	lexer := syntax.InitLexer(source)
	tokens := lexer.Run()
	if lexer.Errors() != nil {
		fmt.Println("----------Errors----------")
		for _, err := range lexer.Errors() {
			fmt.Println(err)
		}
		fmt.Println("--------------------------")
	}
	for _, token := range tokens {
		fmt.Println(token)
	}
}
