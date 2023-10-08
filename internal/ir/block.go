package ir

import (
	"fmt"
	"strings"
)

type Block struct {
	Name string
	Code []*Instruction
}

func NewBlock(name string) *Block {
	return &Block{
		Name: name,
	}
}

func (b *Block) AddInstructions(instrs ...*Instruction) *Block {
	b.Code = append(b.Code, instrs...)
	return b
}

func (b *Block) String() string {
	var buf strings.Builder
	if b.Code != nil {
		if b.Name != "" {
			fmt.Fprintf(&buf, "%s:\n", b.Name)
		}
		for _, instr := range b.Code {
			fmt.Fprintf(&buf, "\t%s\n", instr.String())
		}
	}
	return buf.String()
}
