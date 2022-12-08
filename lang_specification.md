# EBNF
```ebnf
atom        = number 
            | string 
            | "true" 
            | "false" 
            | "nil" 
            | symbol 
            | list 
            | vector
            | hashmap
            | empty;

list        = "'(" {atom} ")" ; 
vector      = "[" {atom} "]" ;
hashmap     = "{" {string|number atom} "}" ;
prefix-op   = add|sub|mul|div|and|or ;
unary-op    = not|neg ;
bin-op      = lt|gt|leq|geq ;
s-expr      = "(" prefix-op {atom|s-expr} ")" 
            | "(" bin-op atom|s-expr atom|s-expr")"
            | "(" unary-op atom|s-expr ")"
            | "(" symbol {atom|s-expr} ")" 
            | "(" do {s-expr} ")" 
            | "(" def symbol atom^empty|s-expr ")"
            | "(" fn vector [string] s-expr ")" ;
```
# Syntax

```clj
(package main)

(import "helloWorld/string" as hlw)
(import "math")

; main function
(fn main [] (
    ; call function from imported package
    (fmt/PrintLn hlw/HellowWorld) ; -> "Hello World"
    ; define varible
    (def pi 3.1415) ; bind 3.1415(float) to pi
    (def r 10) ; bind 10(int) to r
    (def sqrt (* 2 pi (matp/pow r 2))) ; 2 * pi * r * r = 628.32 to sqrt
    (printf "\t%f\n" sqrt) ; "628.32"

    (def res ((fn [a  b ] ( ; lambda function
        (def temp (a + b))
        (ret temp) ; return value from function
    )) 12 10)) ; call function and assign result to varible res

    (def testList [1 2 3 4 5]) ; define list
))

(main) ; call main

```