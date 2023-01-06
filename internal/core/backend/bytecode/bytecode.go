package bytecode

import "fmt"

type Bytecode struct {
	opcode Opcode
	args   any
}

func NewBytecode(opcode Opcode, args any) *Bytecode {
	return &Bytecode{opcode: opcode, args: args}
}

func (b *Bytecode) GetOpcode() Opcode {
	return b.opcode
}

func (b *Bytecode) GetwArgs() any {
	return b.args
}

func (b *Bytecode) GetWidth() int {
	switch args := b.args.(type) {
	case nil:
		return 0
	case string, float64, bool:
		return 1
	case map[string]any:
		return 1
	default:
		panic("unsupported args type: " + fmt.Sprintf("%T", args))

	}
}

func (b *Bytecode) String() string {
	ret := ""
	switch b.opcode {
	case OP_LOAD_CONST:
		ret = "ldc $c,"
	case OP_LOAD_NAME:
		ret = "ldn $s,"
	case OP_SAVE_NAME:
		ret = "svn $s,"
	case OP_SET_NAME:
		ret = "stn $s,"
	case OP_CALL:
		ret = "call $s"
	case OP_SAVE_FN:
		ret = "svf $s,"
	case OP_RET:
		ret = "ret $c,"
	case OP_HLT:
		ret = "hlt,"
	default:
		panic("unecpected opcode: " + fmt.Sprint(b.opcode))
	}
	return ret
}

type Opcode uint8

const (
	OP_NULL Opcode = iota

	OP_LOAD_CONST
	OP_LOAD_NAME
	OP_SAVE_NAME
	OP_SET_NAME

	OP_CALL
	OP_SAVE_FN
	OP_RET

	OP_HLT
)
