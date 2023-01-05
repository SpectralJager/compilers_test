package main

import (
	"grimlang/internal/core/backend/bytecode"
	"grimlang/internal/core/backend/chunk"
	"grimlang/internal/core/backend/object"
)

func main() {
	mainChunk := chunk.NewChunk("main", 0)
	mainChunk.WriteCode(bytecode.OP_LOAD_CONST)
	ind := mainChunk.WriteConst(*object.NewObject(object.ObjectNumber, object.EncodeData(64.)))
	mainChunk.WriteCode(ind)
	mainChunk.WriteCode(bytecode.OP_SAVE_NAME)
	ind = mainChunk.WriteSymbol("a")
	mainChunk.WriteCode(ind)
	mainChunk.Disassebmly()
}
