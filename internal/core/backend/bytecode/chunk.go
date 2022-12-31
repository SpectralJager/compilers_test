package bytecode

import (
	"fmt"
	"grimlang/internal/core/backend/object"
)

type Value struct {
	Type   object.ObjectType
	Object []byte
}

const (
	OP_RETURN byte = iota
	OP_CONSTANT
	OP_NEG
	OP_ADD
	OP_SUB
	OP_DIV
	OP_MUL
)

type Chunk struct {
	Name string
	Code []byte
	Data []Value
}

func (c *Chunk) InitChunk(name string) {
	c.FreeChunk()
	c.FreeData()
	c.Name = name
}

func (c *Chunk) FreeChunk() {
	c.Code = make([]byte, 0)
}

func (c *Chunk) WriteChunk(bt byte) {
	c.Code = append(c.Code, bt)
}

func (c *Chunk) FreeData() {
	c.Data = make([]Value, 0)
}

func (c *Chunk) WriteData(data Value) byte {
	c.Data = append(c.Data, data)
	return byte(len(c.Data) - 1)
}

func (c *Chunk) DisassemblingChunk() {
	fmt.Printf("----%s----\n", c.Name)
	for offset := 0; offset < len(c.Code); offset++ {
		fmt.Printf("%04x\t", offset)
		instr := c.Code[offset]
		switch instr {
		case OP_RETURN:
			fmt.Printf("ret,\n")
		case OP_CONSTANT:
			offset += 1
			fmt.Printf("load_const $%d, // %v\n", c.Code[offset], c.Data[c.Code[offset]])
		case OP_NEG:
			fmt.Printf("neg,\n")
		case OP_ADD:
			fmt.Printf("add,\n")
		case OP_SUB:
			fmt.Printf("sub,\n")
		case OP_MUL:
			fmt.Printf("mul,\n ")
		case OP_DIV:
			fmt.Printf("div,\n ")
		default:
			fmt.Printf("Undefined instruction %x\n", instr)
		}
	}
	fmt.Println()
}
