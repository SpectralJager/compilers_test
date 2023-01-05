package bytecode

const (
	OP_NULL uint8 = iota
	OP_LOAD_CONST
	OP_LOAD_NAME
	OP_SAVE_NAME

	OP_JUMP
	OP_CALL
	OP_RET

	OP_HLT
)
