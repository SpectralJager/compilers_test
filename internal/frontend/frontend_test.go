package frontend

import (
	"encoding/json"
	"testing"
)

func TestLexer(t *testing.T) {
	code := `
	(fn main:int []
		 "help doc"
		(const bazz:int 32)
		(let i:int (imul bazz 10))
		(let j:double 26)
		(let str:string "some string")
		(let list:[]int [1 2 3])
		(let map:{}int {"1"::1 "2"::2})
		(let list_of_maps:[]{}int [{"1"::1} {"1"::1 "2"::2}])
		(println str)
		(println (string (iadd bazz i)))
		(cond 
			((lt i 10)
				(iadd i 10))
			((eq i 10)
				(println (string i)))
			(sub i 10))
		(dotimes l 10
			(println (string l))
			(iadd i l))
		(while (lt i 0)
			(println (string i))
			(isub i 1))
		((lambda [x:int] (println (string x)))
			i)
		(let isum:function (lambda [values:[]int] 
			(let res:int 0)
			(dotimes i (len values)
				(iadd res (lget values i)))
			(ret res)))
		(ret 0)			
	)

	(const fuzz:int 32)
	(var bazz:int 32)
	`
	res, err := Parser.ParseString("",
		code,
		// participle.Trace(os.Stdout),
	)
	if err != nil {
		t.Fatalf("%s", err)
	}
	str, _ := json.MarshalIndent(res, "", " ")
	t.Fatalf("\n%s\n\n%s", Parser.String(), str)
}
