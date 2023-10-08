package ir

import (
	"fmt"
	"log"
	"strings"
)

type Chunk struct {
	Name   string
	Blocks []*Block
}

func NewChunk(name string) *Chunk {
	return &Chunk{
		Name: name,
	}
}

func (ch *Chunk) AddBlocks(blocks ...*Block) *Chunk {
	for _, b := range blocks {
		ch.AddBlock(b)
	}
	return ch
}

func (ch *Chunk) AddBlock(block *Block) *Chunk {
	if bl := ch.GetBlock(block.Name); bl != nil {
		log.Fatalf("Block %v already exists", block.Name)
	}
	ch.Blocks = append(ch.Blocks, block)
	return ch
}

func (ch *Chunk) GetBlock(name string) *Block {
	for _, b := range ch.Blocks {
		if b.Name == name {
			return b
		}
	}
	return nil
}

func (ch *Chunk) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "\n.%s =>\n", ch.Name)
	if len(ch.Blocks) == 0 {
		fmt.Fprint(&buf, "\tempty\n")
	} else {
		for _, bl := range ch.Blocks {
			fmt.Fprintf(&buf, "%s", bl.String())
		}
	}
	return buf.String()
}
