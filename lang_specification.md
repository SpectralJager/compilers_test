# EBNF


# Syntax
```lisp
// simple hello world programm
=>(println "Hello world")
"Hello world"
```
## keywords
- *add* - summation of numbers | concatination of strings | append for vectors | add key value pair to hashmap
```lisp
=>(add 1 5) // sum numbers
6

=>(add "Hello" " " "World") // concat strings
"Hello World"

=>(add [1 2 3] 5 [3 4 5]) // append elements to vector
[1 2 3 5 [3 4 5]]

=>(add {:key1 value1 :key2 value2} :key3 value3)// add key value to hashmap
{:key1 value1 :key2 value2 :key3 value3}

```
- *sub* - subtraction of numbers | remove first math string from source string | remove first math item from vector | remove first math of hashmap **value**
```lisp
=>(sub 3 5)
-2

=>(sub "source string" "source")
" string"

=>(sub [1 2 3] 2)
[1 3]

=>(sub {:key1 val1 :key2 val2} val2)
{:key1 val1}
```
- *div* - divide numbers | remove all math strings from source string | remove all math from vector | remove all math **values** from hashmap
```lisp
=>(div 3 4)
0.75

=>(div "sou2rce str2ing2 " "2")
"source string"

=>(div [1 2 2 3 4 2] 2)
[1 3 4]

=>(div {:key1 val1 :key2 val1 :key3 val2} val1)
{:key3 val2}
```

- *mul* - multiplication of numbers | repeat string n times and concat | multiply each item of vector to n  
```lisp
=>(mul 3 5)
15

=>(mul "str1" 2)
"str1str1"

=>(mul [1 2 3] 2)
[2 4 6]

```

- *def* - bind to a symbol some **atom**
```lisp
=>(def pi 3.14)
@pi: 3.14

=>(def some_list '(add 3 (mul 3 5)))
@some_list: '(add 3 (mul 3 5))
```

- *mov* - assign to the symbol a new value **of the same type**
```lisp
=>(def some_number 123)
@some_number: 123

=>(mov some_number 345)
@some_number: 345

=>(mov some_number "str")
Error: Cannot assign to the some_number of type INT new value of type STRING
```

- *fn* - lambda function
```lisp
=>(def pow2 (fn [x] (mul x x)))
@pow2: fn [x]

=>(pow2 4)
16
```

- *lt* -  less then 
```lisp
=>(lt 3 4)
true

=>(lt 4 3)
false
```

- *gt* -  greater then 
```lisp
=>(gt 3 4)
false

=>(gt 4 3)
true
```

- *le* -  less then or equal 
```lisp
=>(le 3 4)
true

=>(le 4 3)
false

=>(le 4 4)
true
```

- *ge* -  greater then or equal 
```lisp
=>(ge 4 3)
true

=>(ge 3 4)
false

=>(ge 4 4)
true
```

    



```lisp
```
## build-ins