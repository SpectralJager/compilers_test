## Code file
```
@import {
  "std" as std
}

@const {
  a:Int = 1
}

@var {
  str:String = "hello world"
  list:List<String> = ["hello" "world"]
  map:Map<String, String> = {"hello"::"world"}
}

@enum errors:Error {
  ErrNoError => Error("")
  ErrNotImplemented => Error("not implemented")
}

@struct Struct {
  a:Int
  b:String
  fn getA:<Int, errors>() {
    return this.a, ErrNoError
  }
}

@interface IO {
  Read:Int(String)
  Write:Int(String)
}

@fn main:Int() {
  @return 0, ner
}
```

## Fragment file

```
<code>
</code>

<ui>
</ui>
```