package ir

import (
	"fmt"
	"grimlang/internal/frontend"
	"log"
	"reflect"
)

type Chunk interface {
	Meta() string
}

type PackageChunk struct {
	Name      string
	Variables frontend.PackageVariables
	Chunks    []Chunk
}

func (m PackageChunk) Meta() string {
	temp := fmt.Sprintf("Package name: %s\nVariables:\n", m.Name)
	for _, v := range m.Variables.Variables {
		temp += fmt.Sprintf("\t%s %s %v\n", v.Symbol.Name, v.Symbol.Type, v.Value.AtomValue())
	}
	temp += "Chunks:\n"
	for _, ch := range m.Chunks {
		temp += fmt.Sprintf("%s\n", ch.Meta())
	}
	return temp
}

func NewPackageChunk(p *frontend.Package) *PackageChunk {
	ch := PackageChunk{
		Name:      p.Name,
		Variables: p.Variables,
	}
	for _, v := range p.Body {
		switch v := v.(type) {
		case *frontend.FunctionCommand:
			ch.Chunks = append(ch.Chunks, NewFunctionChunk(v))
		default:
			log.Fatalf("unknown node: %v", v)
		}
	}
	return &ch
}

type FunctionChunk struct {
	Symbol frontend.SymbolDeclaration
	Args   []frontend.SymbolDeclaration
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
		temp += fmt.Sprintf("\t%s\n", instr.Decode())
	}
	return temp
}

func NewFunctionChunk(p *frontend.FunctionCommand) *FunctionChunk {
	ch := FunctionChunk{
		Symbol: p.Symbol,
		Args:   p.Args,
	}
	for _, v := range p.Args {
		genInstructions(&ch.Body, v)
	}
	for _, v := range p.Body {
		genInstructions(&ch.Body, v)
	}
	return &ch
}

func genInstructions(instrs *[]Instruction, node any) {
	switch node := node.(type) {
	case frontend.Atom:
		*instrs = append(*instrs, Instruction{
			Op:  OP_LOAD,
			Arg: node.AtomValue(),
			Meta: map[string]string{
				"dataType": node.AtomType(),
			},
		})
	case frontend.SymbolDeclaration:
		*instrs = append(*instrs, Instruction{
			Op:  OP_SAVE,
			Arg: node.Name,
			Meta: map[string]string{
				"symbolType": node.Type,
			},
		})
	case *frontend.LetCommand:
		genInstructions(instrs, node.Value)
		*instrs = append(*instrs, Instruction{
			Op:  OP_SAVE,
			Arg: node.Symbol.Name,
			Meta: map[string]string{
				"symbolType": node.Symbol.Type,
			},
		})
	case *frontend.SetCommand:
		genInstructions(instrs, node.Value)
		*instrs = append(*instrs, Instruction{
			Op:  OP_SAVE,
			Arg: node.Symbol.AtomValue(),
			Meta: map[string]string{
				"symbolType": node.Symbol.AtomType(),
			},
		})
	case *frontend.ReturnCommand:
		genInstructions(instrs, node.ReturnValue)
		*instrs = append(*instrs, Instruction{
			Op: OP_RET,
		})
	case *frontend.IfCommand:
		var ThenInstr, ElseInstr, ElIfInstr []Instruction
		ElIfTemp := make([][]Instruction, len(node.ElseIfBodies))

		genInstructions(instrs, &node.Condition)
		for _, v := range node.ThenBody {
			genInstructions(&ThenInstr, v)
		}
		for i, v := range node.ElseIfBodies {
			genInstructions(&ElIfTemp[i], &node.Condition)
			var tempInstr []Instruction
			for _, j := range v.ThenBody {
				genInstructions(&tempInstr, j)
			}
			ElIfTemp[i] = append(ElIfTemp[i], Instruction{
				Op:  OP_JMC,
				Arg: fmt.Sprint(len(tempInstr) + 1),
			})
			ElIfTemp[i] = append(ElIfTemp[i], tempInstr...)

		}
		for _, v := range node.ElseBodies {
			genInstructions(&ElseInstr, v)
		}

		for i, v := range ElIfTemp {
			ElIfInstr = append(ElIfInstr, v...)
			recLen := 0
			for l, v2 := range ElIfTemp[i:] {
				if l == 0 {
					continue
				}
				recLen += len(v2)
			}
			ElIfInstr = append(ElIfInstr, Instruction{
				Op:  OP_JMP,
				Arg: fmt.Sprint(len(ElseInstr) + recLen),
			})
		}

		ThenInstr = append(ThenInstr, Instruction{
			Op:  OP_JMP,
			Arg: fmt.Sprint(len(ElseInstr) + len(ElIfInstr)),
		})
		*instrs = append(*instrs, Instruction{
			Op:  OP_JMC,
			Arg: fmt.Sprint(len(ThenInstr)),
		})

		*instrs = append(*instrs, ThenInstr...)
		*instrs = append(*instrs, ElIfInstr...)
		*instrs = append(*instrs, ElseInstr...)
	case *frontend.Expression:
		for i := len(node.Args) - 1; i >= 0; i-- {
			genInstructions(instrs, node.Args[i])
		}
		*instrs = append(*instrs, Instruction{
			Op:  OP_CALL,
			Arg: node.Symbol,
		})
	default:
		log.Fatalf("unknown node: %s", reflect.TypeOf(node))
	}
}
