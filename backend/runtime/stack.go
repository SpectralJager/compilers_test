package runtime

import (
	"errors"
	"fmt"
	"grimlang/backend/asm"
	"strings"
)

const (
	StackSize = 1 << 10
)

var (
	ErrStackOverflow  = errors.New("stack overflow")
	ErrStackUnderflow = errors.New("stack underflow")
)

type Stack struct {
	Memory [StackSize]asm.Value
	Sp     int
}

func NewStack() *Stack {
	return &Stack{
		Memory: [1024]asm.Value{},
	}
}

func (st *Stack) Push(val asm.Value) error {
	if st.Sp >= StackSize {
		return ErrStackOverflow
	}
	st.Memory[st.Sp] = val
	st.Sp++
	return nil
}

func (st *Stack) Pop() (asm.Value, error) {
	if st.Sp <= 0 {
		return asm.Value{}, ErrStackUnderflow
	}
	st.Sp--
	return st.Memory[st.Sp], nil
}

func (st *Stack) TraceMemory() string {
	mem := []string{}
	for i := 0; i < st.Sp; i++ {
		val := st.Memory[i]
		mem = append(mem, fmt.Sprintf("#%04x  (%s)%s\n", i, val.Type.Inspect(), val.Inspect()))
	}
	return fmt.Sprintf("=====  Stack trace\n%s", strings.Join(mem, ""))
}