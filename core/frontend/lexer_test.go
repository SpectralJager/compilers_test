package frontend

import (
	"os"
	"testing"
)

func TestLexer(t *testing.T) {
	code := `
@const pi:float = 3.14;

@fn main:void() {
	@var x:int = 2;
	@set x = (floatToInt (add (intToFloat x) pi))
	@if (eq x 10) {
		(println (std/intToString x))
	}
}
`
	lexer := NewLexer(code)
	lexer.Lex()
	if len(lexer.errors) != 0 {
		for _, e := range lexer.errors {
			t.Logf("Error: %v", e)
		}
		t.FailNow()
	}
	// for _, tok := range *tokens {
	// 	t.Logf("%s\n", tok.String())
	// }
	// t.FailNow()
}

func BenchmarkLexer100(b *testing.B) {
	data, _ := os.ReadFile("./test100.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer1000(b *testing.B) {
	data, _ := os.ReadFile("./test1000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer5000(b *testing.B) {
	data, _ := os.ReadFile("./test5000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer10000(b *testing.B) {
	data, _ := os.ReadFile("./test10000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer20000(b *testing.B) {
	data, _ := os.ReadFile("./test20000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer40000(b *testing.B) {
	data, _ := os.ReadFile("./test40000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}
func BenchmarkLexer80000(b *testing.B) {
	data, _ := os.ReadFile("./test80000.txt")
	dataStr := string(data)
	for i := 0; i < b.N; i++ {
		_lexerRun(dataStr, b)
	}
}

func _lexerRun(data string, b *testing.B) {
	lexer := NewLexer(data)
	b.ResetTimer()
	b.ReportAllocs()
	lexer.Lex()
}
