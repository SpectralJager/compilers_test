# Language specification

Basic types:

- strings 
- integers
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