package frontend

import (
	"testing"
)

func TestLexer(t *testing.T) {
	code := `
	(fn Greeting:void [name:string]
		"document string"
		(begin
			(println name)
			nil)
	)
	(fn test:i32 [] 0)

	(fn main:i32 [] 
		(begin
			(let a:i32 10)
			(let b:bool false)
			(let c:f32 12.2)
			(let d:string "some string")
			(set a (iadd a 12))


			0)		
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
