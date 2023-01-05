package bytecode

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"grimlang/internal/core/backend/object"
)

type Value struct {
	Type   object.ObjectType
	Object []byte
}

func (v *Value) Compare(b []byte) bool {
	b = append(b, v.Object...)
	c := 0
	for _, x := range b {
		c ^= int(x)
	}
	return c == 0
}

const (
	OP_RETURN byte = iota
	OP_CONSTANT
	OP_NEG
	OP_ADD
	OP_SUB
	OP_DIV
	OP_MUL
	OP_NOT
	OP_AND
	OP_OR
	OP_LT
	OP_GT
	OP_EQ
	OP_GEQ
	OP_LEQ
	OP_CONCAT
	OP_LEN
	OP_DEF
	OP_SET
	OP_PRINT
)

type Chunk struct {
	Name    string
	Code    []byte
	Data    []Value
	Symbols map[string]byte
}

func (c *Chunk) InitChunk(name string) {
	c.FreeChunk()
	c.FreeData()
	c.FreeSymbols()
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
	for i, val := range c.Data {
		if val.Compare(data.Object) {
			return byte(i)
		}
	}
	c.Data = append(c.Data, data)
	return byte(len(c.Data) - 1)
}

func (c *Chunk) FreeSymbols() {
	c.Symbols = make(map[string]byte)
}

func (c *Chunk) WriteSymbol(symb string, pt byte) byte {
	c.Symbols[symb] = pt
	return byte(len(c.Symbols) - 1)
}
func (c *Chunk) ptToSymb(pt byte) string {
	i := 0
	for k := range c.Symbols {
		if byte(i) == pt {
			return k
		}
		i++
	}
	return ""
}

func (c *Chunk) DisassemblingChunk() {
	fmt.Printf("----%s----\n", c.Name)
	for offset := 0; offset < len(c.Code); offset++ {
		fmt.Printf("%04x \t", offset)
		instr := c.Code[offset]
		switch instr {
		case OP_RETURN:
			fmt.Printf("ret,\n")
		case OP_CONSTANT:
			offset += 1
			fmt.Printf("load_const $%d, // %s\n", c.Code[offset], c.getValue(c.Code[offset]))
		case OP_NEG:
			fmt.Printf("neg, \n")
		case OP_LT:
			fmt.Printf("lt,\n")
		case OP_GT:
			fmt.Printf("gt,\n")
		case OP_EQ:
			fmt.Printf("eq,\n")
		case OP_GEQ:
			fmt.Printf("geq,\n")
		case OP_LEQ:
			fmt.Printf("leq,\n")
		case OP_ADD:
			fmt.Printf("add,\n")
		case OP_SUB:
			fmt.Printf("sub,\n")
		case OP_MUL:
			fmt.Printf("mul,\n")
		case OP_DIV:
			fmt.Printf("div,\n")
		case OP_NOT:
			fmt.Printf("not,\n")
		case OP_AND:
			fmt.Printf("and,\n")
		case OP_OR:
			fmt.Printf("or,\n")
		case OP_LEN:
			fmt.Printf("len,\n")
		case OP_CONCAT:
			fmt.Printf("concat,\n")
		case OP_PRINT:
			fmt.Printf("print,\n")
		case OP_DEF:
			offset += 1
			symb := c.ptToSymb(c.Code[offset])
			val_pt := c.Symbols[symb]
			fmt.Printf("def %s, $%d, // %s = %s\n", symb, val_pt, symb, c.getValue(val_pt))
		default:
			fmt.Printf("Undefined instruction %x\n", instr)
		}
	}
	fmt.Println("\nSymbols:")
	for k, v := range c.Symbols {
		fmt.Printf("%s: %s\n", k, c.getValue(v))
	}
	fmt.Printf("----end of %s----\n", c.Name)
	fmt.Println()
}

func (c *Chunk) getValue(ind byte) string {
	var buf bytes.Buffer
	temp := c.Data[ind]
	_, err := buf.Write(temp.Object)
	if err != nil {
		panic(err)
	}
	switch temp.Type {
	case object.Float:
		return fmt.Sprint(*decode[float64](&buf))
	case object.Bool:
		return fmt.Sprint(*decode[bool](&buf))
	case object.String:
		return fmt.Sprint(*decode[string](&buf))
	default:
		panic("unsupported value type: " + fmt.Sprint(temp.Type))
	}
}

func decode[T any](buf *bytes.Buffer) *T {
	var val T
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&val)
	if err != nil {
		panic(err)
	}
	return &val
}
