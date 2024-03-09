package asm

import (
	"errors"
	"fmt"
	"strings"
)

var (
	fnstr = "(fn %s\n%s%s\n)"

	ErrBlockIndexOutOfBounds = errors.New("block index out of bounds")
)

type Function struct {
	Ident    string
	Blocks   []Block
	Cacheble bool
}

func NewFunction(ident string, cacheble bool, blocks ...Block) Function {
	return Function{
		Ident:    ident,
		Cacheble: cacheble,
		Blocks:   blocks,
	}
}

func (fn *Function) InspectIndent(indent int) string {
	blocks := []string{}
	for i, block := range fn.Blocks {
		blocks = append(blocks, block.InspectIndent(i, indent*2))
	}
	indentStr := strings.Repeat(" ", indent)
	return fmt.Sprintf(fnstr, fn.Ident, indentStr, strings.Join(blocks, "\n"+indentStr))
}

func (fn *Function) Inspect() string {
	return fn.InspectIndent(2)
}

func (fn *Function) Block(bp int) (*Block, error) {
	if bp >= len(fn.Blocks) {
		return &Block{}, ErrBlockIndexOutOfBounds
	}
	return &fn.Blocks[bp], nil
}
