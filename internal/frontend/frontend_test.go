package frontend

import (
	"bytes"
	"testing"

	"github.com/alecthomas/participle/v2"
)

func TestLexer(t *testing.T) {
	code := `
	`

	var buf bytes.Buffer
	_, err := Parser.ParseString("",
		code,
		participle.Trace(&buf),
	)
	if err != nil {
		t.Fatalf("%s\n%s", err, buf.String())
	}
}
