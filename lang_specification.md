# EBNF
```ebnf
atom        = int
            | float
            | string 
            | "true" | "false" 
            | symbol
            | function
            | list 
            | empty;
list        = "'[" {atom} "]" ; 
program     = {s-expr | sp-form | atom} ;
s-expr      = "(" op {s-expr|atom} ")";
op          = symbol | build-ins | function ;
sp-forms    = def | set | fn ;
def         = "(" "def" symbol atom ")" ;
set         = "(" "set" symbol atom ")" ;
ret         = "(" "ret" atom|s-expr ")" ;
fn          = "(" "fn" [symbol] "["{symbol}"]" "("{s-expr|sp-form}")" ")" ;
```
# Syntax

```clj
(package main)

(import "helloWorld/string" as hlw)
(import "math")

; main function
(fn main [] (
    ; call function from imported package
    (printf hlw/HellowWorld) ; -> "Hello World"
    ; define varible
    (def pi 3.1415) ; bind 3.1415(float) to pi
    (def r 10) ; bind 10(int) to r
    (def sqrt (mul 2 pi (matp/pow r 2))) ; 2 * pi * r * r = 628.32 to sqrt
    (printf "\t%f\n" sqrt) ; "628.32"
    (def res ((fn [a  b ] ( ; lambda function
        (def temp (a + b))
        (ret temp) ; return value from function
    )) 12 10)) ; call function and assign result to varible res
    (def testList [1 2 3 4 5]) ; define list
))

(main) ; call main

```