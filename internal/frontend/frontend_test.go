package frontend

import (
	"testing"
)

func TestLexer(t *testing.T) {
	code := `
	; comment
	(fn Greeting:void [name:string]
		"document string"
		(begin
			(let lst:[]i32 [ 2 3 4 5 ])
			(let mp:{}i32 { "1"::1 "2"::3 })
			(println name)
			nil)
	)
	(fn test:i32 [] 0) ; comment

	(glob lala:f32 12.2)
	(const lulu:string "Hello")

	(fn main:void [] ; comment
		(begin ; comment
			(let a:i32 10)
			(let b:bool false)
			(let c:f32 12.2)
			(let d:string "some string")
			(set a (iadd a 12))
			(const e:string "world\"")
			(println (concat lulu e))
			(let x:i32 10)
			(if (leq x 10)
				(println (string x))
				(set x 10))
			(cond 
				((lt x 10) 
					(println (string x)))
				((lt x 10) 
					(begin 
						(let y:i32 256)
						(set x (mul x y))))
				(println "default"))
			nil)		
	)

	`
	_, err := parser.ParseString("",
		code,
		// participle.Trace(os.Stdout),
	)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Fatal("\n" + parser.String())
}
