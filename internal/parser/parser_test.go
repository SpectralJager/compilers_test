package parser

import (
	"encoding/xml"
	"os"
	"strings"
	"testing"

	"github.com/alecthomas/participle/v2"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		desc string
		src  string
	}{
		{
			desc: "parse global variable",
			src:  "../../src/glob_var.grim",
		},
		{
			desc: "parse functions",
			src:  "../../src/func.grim",
		},
	}
	out, err := os.OpenFile("output.xml", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		t.Fatalf("can't open output.xml: %s", err.Error())
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			data, err := os.ReadFile(tC.src)
			if err != nil {
				t.Fatalf("can't read test source file %s: %s", tC.src, err.Error())
			}
			var errBuf strings.Builder
			prog, err := Parser.ParseBytes(tC.src, data, participle.Trace(&errBuf))
			if err != nil {
				t.Log(errBuf.String())
				t.Fatalf("can't parse test source file %s: %s", tC.src, err.Error())
			}
			prog.Name = tC.desc
			tmp, err := xml.MarshalIndent(prog, "", "\t")
			if err != nil {
				t.Fatalf("can't marshal program: %s", err.Error())
			}
			out.Write(tmp)
			out.Write([]byte{'\n'})
			t.Log("\n" + prog.String())
		})
	}
}
