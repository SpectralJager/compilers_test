package main

import (
	"fmt"
	"gl/internal/repr"
	"gl/internal/vm"
)

func main() {
	ch := repr.NewChunk("test")
	ch.WriteBytes(
		repr.OP_HALT,
	)
	fmt.Println(ch.Disassembly())
	v := vm.NewVM()
	v.ExecuteChunk(ch)
	fmt.Println(v.StackTrace())
}
