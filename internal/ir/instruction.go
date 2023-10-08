package ir

type Instruction struct {
	Operation OperationType
	Args      []InstrArg
}

type OperationType string

const (
	OP_NULL OperationType = "null"

	OP_GLOBAL_LOAD = "global.load"
	OP_GLOBAL_SAVE = "global.save"
)
