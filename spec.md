# Language specification

Basic types:

- strings 
- integers
- boolean
- float

Data types:

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
  - `'@if' exprArg '{' locals+ '}' ('elif' exprArg '{' locals+ '}')* ('else' '{' locals+ '}')?`
- @for: 
  - `'@for' SYMBOL ':' type 'from' atom 'to' atom '{' locals+ '}'`
- @while: 
  - `'@while' expr '{' locals+ '}' ('else' '{' locals+ '}')?`
