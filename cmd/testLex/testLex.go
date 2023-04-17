package main

import (
	"fmt"
	"gl/src/syntax"
)

func main() {
	source := `
@import "std" as std;
@import "http" as http;
@import "db" as db;

const a:i32 = 1;

pub fn main() void {
	@var server: http/Server = (http/NewServer);
	@set server/address = "localhost";
	@set server/port = 8080;

	@var index:http/Handler = (http/NewHandler "/" fn(ctx: http/RequestContext) http/Response {
		@if (eq ctx/Method "GET") {
			@return http/Response{status::200 body::"Hello World!"};
		} @elif (eq ctx/Method "POST") {
			@var name: string = ctx/Params["name"];
			@return http/Response{status::200 body::"Hello " + name + "!"};
		} @else {
			@return http/Response{status::404 body::"Not Found"};
		}
	});

	(server/serve)
}

fn infoHandler(ctx: http/RequestContext) http/Response {
	@var users: List<db/User> = (db/GetUsersAll);
	@each user in users {
		@if (eq ctx/Params["name"] user/name){
			@return http/Response{status::200 body::(user/ToJson)};
		} @else {
			@return http/Response{status::403 body::"User not found"};
		}
	}
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
