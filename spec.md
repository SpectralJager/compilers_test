```
// function stmts
@fn main(x:int y:float) <int> {...}

// varible stmts
@var x:int = 1;
@var x:float = 1.0;
@var x:string = "hello world";
@var x:bool = false;
@var y:list<int> = '(1 2 3);
@var y:map<string int> = #("string" -> 1);

// expressions
(function arg1 arg2)

// condition stmts
@if (leq a b) {
  ...
}
@if (leq a b) {
  ...
} else {
  ...
}
@if (leq a b) {
  ...
} (neq a b) {
  ...
} else {
  ...
}

// loop stmts
@while (eq a b) {
  ...
}
@each val:int <- '(1 2 3) {
  ...
}

// other stmts
@return ...;
@set a = 1;
```