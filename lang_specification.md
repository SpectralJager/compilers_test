# EBNF
```ebnf
atom        = int
            | float
            | string 
            | "true" | "false" 
            | symbol
            | empty;
program     = {expr} ;
expr        = s-expr | sp-form ;
s-expr      = "(" symbol {s-expr|atom} ")";
sp-forms    = def | set | fn | ret;
def         = "(" "def" symbol atom|s-expr ")" ;
set         = "(" "set" symbol atom|s-expr ")" ;
ret         = "(" "ret" atom ")" ;
fn          = "(" "fn" symbol "("{symbol}")" "("{expr}")" ")" ;
```
# Syntax

```clj
(package main)


; main function
(fn main [] (
    ; call function from imported package
    (printf "hello world") ; -> "Hello World"
    ; define varible
    (def pi 3.1415) ; bind 3.1415(float) to pi
    (def r 10) ; bind 10(int) to r
    (def sqrt (mul 2 pi (pow r 2))) ; 2 * pi * r * r = 628.32 to sqrt
    (printf "\t%f\n" sqrt) ; "628.32"
))

(main) ; call main

```