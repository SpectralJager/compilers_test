package bytecode

import "fmt"

type Bytecode struct {
	opcode Opcode
	args   interface{}
}

func NewBytecode(opcode Opcode, args interface{}) *Bytecode {
	return &Bytecode{opcode: opcode, args: args}
}

func (b *Bytecode) GetOpcode() Opcode {
	return b.opcode
}
func (b *Bytecode) GetwArgs() interface{} {
	return b.args
}

func (b *Bytecode) GetWidth() int {
	switch args := b.args.(type) {
	case nil:
		return 0
	case string, float64, bool:
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
		ret = "ldn $n,"
	case OP_SAVE_NAME:
		ret = "svn $n, $c,"
	case OP_JUMP:
		ret = "jmp $p,"
	case OP_CALL:
		ret = "call $n,"
	case OP_RET:
		ret = "ret,"
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

	OP_JUMP
	OP_CALL
	OP_RET

	OP_HLT
)
