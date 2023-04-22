# Language specification

Basic types:

- strings 
    - unicode
- integers
    - i8, i16, i32, i64
    - u8, u16, u32, u64
- bool
- float
    - f64

Data types:

- List
- Maps
- Turple
- Enum
- Struct

Keywords:

- @set: 
  - `'@set' SYMBOL '=' expr ';'`
- @var: 
  - `'@var' SYMBOL ':' type '=' expr ';'`
- @const: 
  - `'@const' SYMBOL ':' type '=' atom ';'`
- @import: 
  - `'@import' STRING 'as' SYMBOL ';'`
- @fn: 
  - `'@fn' SYMBOL ':' type '(' (SYMBOL ':' type)* ')' '{' locals+ '}'`
- @lambda: 
  - `'@lambda' type '(' (SYMBOL ':' type)* ')' '{' locals+ '}'`
- @if: 
  - `'@if' expr '{' locals+ '}' ('elif' expr '{' locals+ '}')* ('else' expr '{' locals+ '}')?`
- @for: 
  - `'@for' SYMBOL ':' type 'from' atom 'to' atom '{' locals+ '}'`
- @while: 
  - `'@while' expr '{' locals+ '}'`