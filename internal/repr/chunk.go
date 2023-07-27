package repr

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
-------- Name --------
constants:
	...
code:
	0000 null
	...
	002f hlt
*/

type Chunk struct {
	Name      string
	Code      []byte
	ConstPool []Object
}

func NewChunk(name string) *Chunk {
	return &Chunk{
		Name:      name,
		Code:      make([]uint8, 0, 2048),
		ConstPool: make([]Object, 0, 2048),
	}
}

func (ch *Chunk) WriteConstant(v Object) int {
	ch.ConstPool = append(ch.ConstPool, v)
	return len(ch.ConstPool) - 1
}

func (ch *Chunk) GetConstant(index int) (Object, error) {
	if index >= len(ch.ConstPool) {
		return Object{}, fmt.Errorf("cant find value with index %d", index)
	}
	return ch.ConstPool[index], nil
}

func (ch *Chunk) WriteBytes(bs ...byte) {
	ch.Code = append(ch.Code, bs...)
}

func (ch *Chunk) Disassembly() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "-------- %s --------\n", ch.Name)
	fmt.Fprint(&buf, "==== constants: \n")
	for i, c := range ch.ConstPool {
		fmt.Fprintf(&buf, "$%d %s\n", i, c.String())
	}
	fmt.Fprint(&buf, "==== code: \n")
	for offset := 0; offset < len(ch.Code); {
		intsr := ch.Code[offset]
		switch intsr {
		case OP_NULL:
			fmt.Fprintf(&buf, "%04x %s\n", offset, "null")
			offset += 1
		case OP_HALT:
			fmt.Fprintf(&buf, "%04x %s\n", offset, "halt")
			offset += 1
		case OP_CONST:
			fmt.Fprintf(&buf, "%04x %s $%d\n", offset, "const", ch.Code[offset+1])
			offset += 2
		case OP_INEG:
			fmt.Fprintf(&buf, "%04x %s\n", offset, "ineg")
			offset += 1
		case OP_IINC:
			fmt.Fprintf(&buf, "%04x %s\n", offset, "iinc")
			offset += 1
		case OP_IDEC:
			fmt.Fprintf(&buf, "%04x %s\n", offset, "idec")
			offset += 1
		case OP_IADD:
			fmt.Fprintf(&buf, "%04x %s\n", offset, "iadd")
			offset += 1
		case OP_ISUB:
			fmt.Fprintf(&buf, "%04x %s\n", offset, "isub")
			offset += 1
		case OP_IMUL:
			fmt.Fprintf(&buf, "%04x %s\n", offset, "imul")
			offset += 1
		case OP_IDIV:
			fmt.Fprintf(&buf, "%04x %s\n", offset, "idiv")
			offset += 1
		case OP_JMP:
			fmt.Fprintf(&buf, "%04x %s %04xh\n", offset, "jmp", binary.BigEndian.Uint16(ch.Code[offset+1:offset+3]))
			offset += 3
		case OP_JZ:
			fmt.Fprintf(&buf, "%04x %s %04xh\n", offset, "jz", binary.BigEndian.Uint16(ch.Code[offset+1:offset+3]))
			offset += 3
		case OP_JNZ:
			fmt.Fprintf(&buf, "%04x %s %04xh\n", offset, "jnz", binary.BigEndian.Uint16(ch.Code[offset+1:offset+3]))
			offset += 3
		default:
			panic(fmt.Errorf("unexpected opcode %04xh", intsr))
		}
	}
	return buf.String()
}
