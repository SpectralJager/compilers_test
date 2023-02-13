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
- String - string;
- Any - any;
- Null value - nil;

### Composite types

A composite value is a value that is composed from simpler values:

- List - []*type*;
- Map - {}*type*;
- Struct - struct;

## Keywords

- package - package declaration command
- import - import command
- const - define constant value command
- glob - define global varible command
- let - define local varible inside block command
- fn - function declaration command
- cond - multiple condition expressions
- while - while loop command
- dotimes - dotimes loop command
- ret - return command
- nil 
