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