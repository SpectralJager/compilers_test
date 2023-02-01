# Specification

Grimlang is: 

- general-purpose language, designed for backend development. 
- staticly typed. 


## Build-in types

### Primitive types

A primitive value is one that cannot be decomposed into simpler values:

- Integer - i8, i16, i32, u8, u16, u32;
- Float - f32;
- Bool - bool;
- Any - any;
- Null value - nil;

### Composite types

A composite value is a value that is composed from simpler values:

- List - list<*type*>;
- Map - map<*type*>;
- Struct - struct;
- String - string;

## Keywords

- package - package declaration command
- import - import command
- const - define constant value command
- glob - define global varible command
- let - define local varible inside block command
- fn - function declaration command
- begin - block declaration command
- if - if command
- cond - multiple condition expressions
- while - while loop command
- dotimes - dotimes loop command
- ret - return command
- break - break current loop command
- continue - next iteratoin of current loop command
- nil 

## Examples

```clj
(package main)

; function declaratino
; fn *name*:*return-type* "*docstring, optional*" [*args*:*type*] body
(fn Greeting:void [x:string]
    (begin
        (printf "Hello, %s\n" x)
        nil)
)


(fn main:void [] ; funciton main is entry point of programm
    (begin
        ; declaration of local varible
        (let str:string (concat "Hello" ", " "World!"))
        (println str) ; => Hello, World!
        ; ...............................
        (let x:i32 (sLen str))
        ; if (condiiton) (true-body) (false-body)
        (if (leq x 10) ; condition
            (println  (string x)) ; true
            (println "more then 10")) ; false

        ; cond ((condition) (true-body))... (else-body)
        (cond 
            ; if x < 10 then println x
            ((lt x 10) (println (string x))) 
            ; else if x < 15 then println x
            ((lt x 15) (println "less then 15")) 
            ; else
            (println "more then 10")) 

        ; dotimes symb:i32 times:i32 (body)
        (dotimes i 5
            (println (string i))) ; => 0, 1, 2, 3, 4
        
        ; while (condition) (true-body)
        (set x 0) ; x = 0
        (while (lt x 5)
            (begin 
                (println (string x)) ; => 0, 1, 2, 3, 4
                (inc x) ; x++
                nil)) 

        ; call function
        (Greeting "Alex") ; => Hello, Alex
        (let name:string "Alex")
        (Greeting name) ; => Hello, Alex
        ; call function with named arg - varible::arg-name
        (Greeting name::x) ; => Hello, Alex

        ; declaration of list
        (let lst:list<i32> [2 32 222 12])
        (iterList item:i32 lst 
            (println (string item))) ; => 2, 32, 222, 12
        (println (string (lGet lst 2))) ; => 222

        ; declaration of map
        (let mp:map<i32> {"1"::2,"2"::32,"3"::222,"4"::12})
        (iterMap key:string item:i32 mp 
            (printf "%s-%d" key item)) ; => 1-2, 2-32, 3-222, 4-12
        (println (string (mGet mp "1"))) ; => 2

        nil)
)
```
