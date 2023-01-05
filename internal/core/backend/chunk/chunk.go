package chunk

import (
	"fmt"
	"grimlang/internal/core/backend/bytecode"
	"grimlang/internal/core/backend/object"
)

type Chunk struct {
	Name         string
	StartAddress uint16
	Code         []uint8
	Data         []object.Object
	Env          []string
}

func NewChunk(name string, sa uint16) *Chunk {
	return &Chunk{Name: name, StartAddress: sa}
}

func (c *Chunk) WriteCode(code uint8) {
	c.Code = append(c.Code, code)
}

func (c *Chunk) WriteConst(con object.Object) uint8 {
	c.Data = append(c.Data, con)
	return uint8(len(c.Data) - 1)
}

func (c *Chunk) WriteSymbol(symb string) uint8 {
	c.Env = append(c.Env, symb)
	return uint8(len(c.Env) - 1)
}

func (c *Chunk) Disassebmly() {
	fmt.Printf("----------%s----------\n", c.Name)
	for i := 0; i < len(c.Code); i++ {
		offset := c.StartAddress + uint16(i)
		fmt.Printf("%04x\t", offset)
		switch c.Code[i] {
		case bytecode.OP_LOAD_CONST:
			i += 1
			fmt.Printf("ldc $c,\n")
		case bytecode.OP_SAVE_NAME:
			i += 1
			fmt.Printf("def $s,\n")
		default:
			panic("unsupported bytecode operation: " + fmt.Sprint(c.Code[i]))
		}
	}
	fmt.Printf("\nSymbols:\n")
	for i, v := range c.Env {
		fmt.Printf("$%d: %s\n", i, v)
	}
	fmt.Printf("------End of %s-------\n\n", c.Name)
}
