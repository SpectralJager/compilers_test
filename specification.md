# Specification
## Base types
- int
- float
- string
- bool
- list<Type>
- hashmap<Type>
- enum
- null<type>
## User defined types
- struct {Field:Type ...}
## Language constructions
### Memory manipulation
- declare variable -> @var name:type = expression;
- set variable -> @set name = expression;
### Flow constructions
- if-elif-else -> @if condition {
    body...
elif condition =>
    body...
else =>
    body...
}
- for -> @for i = 0 to number {
    body...
else =>
    body...
}
- while -> @while condition {
    body...
else =>
    body...
}
- for-each -> @foreach item in someList {
    body...
else =>
    body...
}
- return -> @ret expression;
- break -> @break;
- continue -> @continue;
### Functions
- definition -> @fn name[arg:type ...] <Type> {body...}
- call -> (name args...)
### Error 
- Error rising -> @raise message;
- Error handling -> @catch expression => |msg| (println msg);
## Examples
```

```