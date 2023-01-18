package chunk

import (
	"fmt"
	"grimlang/internal/core/backend/v1/bytecode"
)

type Chunk struct {
	Name        string
	Bytecodes   []bytecode.Bytecode
	StartOffset uint16
}

func NewChunk(name string, startOffset uint16) *Chunk {
	return &Chunk{Name: name, StartOffset: startOffset}
}

func (c *Chunk) WriteBytecode(bt bytecode.Bytecode) {
	c.Bytecodes = append(c.Bytecodes, bt)
}

func (c *Chunk) Disassembly() {
	fmt.Printf("\n--------%s--------\n", c.Name)
	offset := c.StartOffset
	for _, bt := range c.Bytecodes {
		fmt.Printf("%04x\t%s\n", offset, bt.String())
		offset += uint16(bt.GetWidth() + 1)
	}
	fmt.Printf("-----end of %s----\n\n", c.Name)
}
