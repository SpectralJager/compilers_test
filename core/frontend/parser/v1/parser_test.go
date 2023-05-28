package parser

// func TestParser(t *testing.T) {
// 	code := `
// 	@const a:int = 1;
// 	@var b:int = (add a 2);

// 	@fn sum:void(a:int b:int) {
// 		@var res:int = (add a b);
// 		@set res = (add res 2);
// 		@if (eq res 10) {
// 			(printf "res = 10")
// 		} elif (eq res 15) {
// 			(printf "res = 15")
// 		} elif (eq res 20) {
// 			(printf "res = 20")
// 		} else {
// 			(printf "res = %d" res)
// 		}

// 		@while false {
// 			(printf "while")
// 		} else {
// 			(printf "else")
// 		}

// 		@for i from 0 to 10 {
// 			(printf "%d" i)
// 		}
// 	}
// 	`
// 	lex := lexer.NewLexer(code)
// 	tokens := lex.Lex()
// 	if len(lex.Error()) != 0 {
// 		for _, e := range lex.Error() {
// 			t.Logf("Error: %v", e)
// 		}
// 		t.FailNow()
// 	}
// 	prs := NewParser(*tokens)
// 	programm := prs.Parse()
// 	if len(prs.Errors()) != 0 {
// 		for _, e := range prs.Errors() {
// 			t.Logf("Error: %v", e)
// 		}
// 		t.FailNow()
// 	}
// 	data, _ := json.MarshalIndent(programm, "", "\t")
// 	t.Fatalf("%s\n", data)
// }
