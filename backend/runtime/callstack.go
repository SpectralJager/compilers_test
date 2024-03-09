package runtime

import (
	"fmt"
	"strings"
)

const (
	CallStackSize = 1 << 9
)

type CallStack struct {
	Calls [CallStackSize]Frame
	Sp    int
}

func NewCallStack() *CallStack {
	return &CallStack{
		Calls: [CallStackSize]Frame{},
	}
}

func (st *CallStack) Push(fr Frame) {
	if st.Sp >= CallStackSize {
		panic(ErrStackOverflow)
	}
	st.Calls[st.Sp] = fr
	st.Sp++
}

func (st *CallStack) Pop() {
	if st.Sp <= 0 {
		panic(ErrStackUnderflow)
	}
	st.Sp--
}

func (st *CallStack) Top() *Frame {
	return &st.Calls[st.Sp-1]
}

func (st *CallStack) TraceCalls() string {
	calls := []string{}
	for i := 0; i < st.Sp; i++ {
		fr := st.Calls[i]
		calls = append(calls, fmt.Sprintf("#%04x  %s\n", i, fr.Function.Ident))
	}
	return fmt.Sprintf("=====  Stack trace\n%s", strings.Join(calls, ""))
}
