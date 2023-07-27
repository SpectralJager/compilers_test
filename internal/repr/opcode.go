package repr

const (
	OP_NULL byte = iota
	OP_HALT
	OP_CONST
	OP_JMP
	OP_JNZ
	OP_JZ

	// integers buildins
	OP_INEG
	OP_IINC
	OP_IDEC
	OP_IADD
	OP_ISUB
	OP_IMUL
	OP_IDIV
)
