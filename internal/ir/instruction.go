package ir

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Code []byte

func NewCode() *Code {
	return &Code{}
}

func (c *Code) WriteByte(b byte) *Code {
	*c = append(*c, b)
	return c
}

func (c *Code) WriteBytes(b ...byte) *Code {
	*c = append(*c, b...)
	return c
}

func (c *Code) ReadBytes(offset, n int) []byte {
	return (*c)[offset : offset+n]
}

func (c *Code) Len() int { return len(*c) }

const (
	OP_HALT byte = iota
	OP_RETURN
	OP_LOAD
	OP_GLOBAL_SET
	OP_GLOBAL_LOAD
	OP_GLOBAL_SAVE
	OP_LOCAL_SET
	OP_LOCAL_LOAD
	OP_LOCAL_SAVE
	OP_CALL
	OP_JUMP
	OP_JUMP_CONDITION
	OP_INT_FUNC
	OP_FLOAT_FUNC
	OP_STRING_FUNC
)

func Halt() byte               { return OP_HALT }
func Return(count byte) []byte { return []byte{OP_RETURN, count} }
func Load(index uint32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_LOAD)
	binary.Write(&buf, binary.LittleEndian, index)
	return buf.Bytes()
}
func GlobalSet(g, c uint32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_GLOBAL_SET)
	binary.Write(&buf, binary.LittleEndian, g)
	binary.Write(&buf, binary.LittleEndian, c)
	return buf.Bytes()
}
func GlobalLoad(g uint32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_GLOBAL_LOAD)
	binary.Write(&buf, binary.LittleEndian, g)
	return buf.Bytes()
}
func GlobalSave(g uint32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_GLOBAL_SAVE)
	binary.Write(&buf, binary.LittleEndian, g)
	return buf.Bytes()
}
func LocalSet(g, c uint32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_LOCAL_SET)
	binary.Write(&buf, binary.LittleEndian, g)
	binary.Write(&buf, binary.LittleEndian, c)
	return buf.Bytes()
}
func LocalLoad(g uint32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_LOCAL_LOAD)
	binary.Write(&buf, binary.LittleEndian, g)
	return buf.Bytes()
}
func LocalSave(g uint32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_LOCAL_SAVE)
	binary.Write(&buf, binary.LittleEndian, g)
	return buf.Bytes()
}
func Call(g uint32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_CALL)
	binary.Write(&buf, binary.LittleEndian, g)
	return buf.Bytes()
}
func Jump(addr uint32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_JUMP)
	binary.Write(&buf, binary.LittleEndian, addr)
	return buf.Bytes()
}
func JumpCondition(addr uint32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_JUMP_CONDITION)
	binary.Write(&buf, binary.LittleEndian, addr)
	return buf.Bytes()
}
func IntFunc(id byte) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_INT_FUNC)
	binary.Write(&buf, binary.LittleEndian, id)
	return buf.Bytes()
}
func FloatFunc(id byte) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_FLOAT_FUNC)
	binary.Write(&buf, binary.LittleEndian, id)
	return buf.Bytes()
}
func StringFunc(id byte) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, OP_STRING_FUNC)
	binary.Write(&buf, binary.LittleEndian, id)
	return buf.Bytes()
}

func (c *Code) Disassembly() string {
	var buf bytes.Buffer
	bts := (*c)[:]
	for ip := 0; ip < len(bts); {
		i := bts[ip]
		switch i {
		case OP_HALT:
			ip += 1
			fmt.Fprintf(&buf, "%08x halt\n", ip-1)
		case OP_RETURN:
			temp := ip
			ip += 1
			count := c.ReadBytes(ip, 1)[0]
			ip += 1
			fmt.Fprintf(&buf, "%08x return %d\n", temp, count)
		case OP_LOAD:
			temp := ip
			ip += 1
			ind := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			fmt.Fprintf(&buf, "%08x load $%d\n", temp, ind)
		case OP_GLOBAL_SET:
			temp := ip
			ip += 1
			indV := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			indC := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			fmt.Fprintf(&buf, "%08x global_set #%d $%d\n", temp, indV, indC)
		case OP_GLOBAL_LOAD:
			temp := ip
			ip += 1
			ind := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			fmt.Fprintf(&buf, "%08x global_load #%d\n", temp, ind)
		case OP_GLOBAL_SAVE:
			temp := ip
			ip += 1
			ind := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			fmt.Fprintf(&buf, "%08x global_save #%d\n", temp, ind)
		case OP_LOCAL_SET:
			temp := ip
			ip += 1
			indV := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			indC := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			fmt.Fprintf(&buf, "%08x local_set #%d $%d\n", temp, indV, indC)
		case OP_LOCAL_LOAD:
			temp := ip
			ip += 1
			ind := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			fmt.Fprintf(&buf, "%08x local_load #%d\n", temp, ind)
		case OP_LOCAL_SAVE:
			temp := ip
			ip += 1
			ind := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			fmt.Fprintf(&buf, "%08x local_save #%d\n", temp, ind)
		case OP_CALL:
			temp := ip
			ip += 1
			ind := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			fmt.Fprintf(&buf, "%08x call #%d\n", temp, ind)
		case OP_JUMP:
			temp := ip
			ip += 1
			addr := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			fmt.Fprintf(&buf, "%08x jump [%08x]\n", temp, addr)
		case OP_JUMP_CONDITION:
			temp := ip
			ip += 1
			addr := binary.LittleEndian.Uint32(c.ReadBytes(ip, 4))
			ip += 4
			fmt.Fprintf(&buf, "%08x jump_condition [%08x]\n", temp, addr)
		case OP_INT_FUNC:
			temp := ip
			ip += 1
			id := c.ReadBytes(ip, 1)[0]
			ip += 1
			fmt.Fprintf(&buf, "%08x int_func %d\n", temp, id)
		case OP_FLOAT_FUNC:
			temp := ip
			ip += 1
			id := c.ReadBytes(ip, 1)[0]
			ip += 1
			fmt.Fprintf(&buf, "%08x float_func %d\n", temp, id)
		case OP_STRING_FUNC:
			temp := ip
			ip += 1
			id := c.ReadBytes(ip, 1)[0]
			ip += 1
			fmt.Fprintf(&buf, "%08x string_func %d\n", temp, id)
		}
	}
	return buf.String()
}
