package asm

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrReachEndOfBlock = errors.New("reached end of block")
	blstr              = "(block %d\n%s%s)"
)

type Block struct {
	Ip           int
	Instructions []Instruction
}

func NewBlock(instrs ...Instruction) *Block {
	return &Block{
		Instructions: instrs,
	}
}

func (bl *Block) Next() (Instruction, error) {
	if bl.Ip >= len(bl.Instructions) {
		return Instruction{}, ErrReachEndOfBlock
	}
	instr := bl.Instructions[bl.Ip]
	bl.Ip++
	return instr, nil
}

func (bl *Block) InspectIndent(index, indent int) string {
	instrs := []string{}
	for _, instr := range bl.Instructions {
		instrs = append(instrs, instr.Inspect())
	}
	indentStr := strings.Repeat(" ", indent)
	return fmt.Sprintf(blstr, index, indentStr, strings.Join(instrs, "\n"+indentStr))
}

func (bl *Block) Inspect(index int) string {
	return bl.InspectIndent(index, 4)
}
