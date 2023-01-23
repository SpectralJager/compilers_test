# Specification

Grimlang is: 

- general-purpose language, designed for backend development. 
- staticly typed. 


## Build-in types

### Primitive types

A primitive value is one that cannot be decomposed into simpler values:

- Integer - i8, i16, i32, ui8, ui16, ui32:
- Float - f32:
- Rune - rune:
- Bool - bool:
- Any - any;

### Composite types

A composite value is a value that is composed from simpler values:

- List - list<*type*>;
- Map - map<*type*>;
- Enum;
- Struct; 
- String;

## Keywords

- const - define constant value
- glob - define global varible
- let - define local varible
- fn - function declaration
- clojure - clojure block declaration
- if - if command
- ret - return command
- foreach - for each loop
- for - for loop

## Project structure

