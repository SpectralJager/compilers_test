package ir

import (
	"fmt"
	"strings"
)

type Chunk struct {
	Name string
	Code []Instruction
}

func NewChunk(name string, code ...Instruction) Chunk {
	return Chunk{
		Name: name,
		Code: code,
	}
}

func (ch *Chunk) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s |>\n", ch.Name)
	for i, instr := range ch.Code {
		fmt.Fprintf(&buf, "%04x %s\n", i, instr.String())
	}
	return buf.String()
}

func (ch *Chunk) AddCode(code ...Instruction) {
	ch.Code = append(ch.Code, code...)
}
