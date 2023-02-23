package frontend

import (
	"fmt"
	"log"
)

type Chunk interface {
	Meta() string
}

type PackageChunk struct {
	Name      string
	Variables PackageVariables
	Chunks    []Chunk
}

func (m PackageChunk) Meta() string {
	temp := fmt.Sprintf("Package name: %s\nVariables:\n", m.Name)
	for _, v := range m.Variables.Variables {
		temp += fmt.Sprintf("\t%s %s %v\n", v.Symbol.Name, v.Symbol.Type, v.Value)
	}
	temp += "Chunks:\n"
	for _, ch := range m.Chunks {
		temp += fmt.Sprintf("%s\n", ch.Meta())
	}
	return temp
}

func NewPackageChunk(p *Package) *PackageChunk {
	ch := PackageChunk{
		Name:      p.Name,
		Variables: p.Variables,
	}
	for _, v := range p.Body {
		switch v := v.(type) {
		case *FunctionCommand:
			ch.Chunks = append(ch.Chunks, NewFunctionChunk(v))
		default:
			log.Fatalf("unknown node: %v", v)
		}
	}
	return &ch
}

type FunctionChunk struct {
	Symbol SymbolDeclaration
	Args   []SymbolDeclaration
	Body   []Instruction
}

func (m FunctionChunk) Meta() string {
	temp := fmt.Sprintf("; %s\n", m.Symbol.Name)
	temp += "; "
	if len(m.Args) > 0 {
		for _, v := range m.Args {
			temp += fmt.Sprintf("%s ", v.Type)
		}
	} else {
		temp += "void "
	}
	temp += fmt.Sprintf("-> %s\n", m.Symbol.Type)
	for _, instr := range m.Body {
		temp += fmt.Sprintf("\t%s\n", instr.inst())
	}
	return temp
}

func NewFunctionChunk(p *FunctionCommand) *FunctionChunk {
	ch := FunctionChunk{
		Symbol: p.Symbol,
		Args:   p.Args,
	}
	for _, v := range p.Body {
		genInstructions(&ch.Body, v)
	}
	return &ch
}

func genInstructions(instrs *[]Instruction, node any) {
	switch node := node.(type) {
	case *Int:
		*instrs = append(*instrs, &INST_LC{node.Value})
	case *String:
		*instrs = append(*instrs, &INST_LC{node.Value})
	case *Float:
		*instrs = append(*instrs, &INST_LC{node.Value})
	case *Symbol:
		*instrs = append(*instrs, &INST_LS{node.Value})
	case *LetCommand:
		genInstructions(instrs, node.Value)
		*instrs = append(*instrs, &INST_SCS{node.Symbol.Name})
	case *Expression:
		for i := len(node.Args) - 1; i >= 0; i-- {
			genInstructions(instrs, node.Args[i])
		}
		*instrs = append(*instrs, &INST_CALL{node.Symbol})
	default:
		log.Fatalf("unknown node: %v", node)
	}
}
