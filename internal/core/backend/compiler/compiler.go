package main

import (
	"bytes"
	"encoding/gob"
	"grimlang/internal/core/backend/bytecode"
	"grimlang/internal/core/backend/object"
	"grimlang/internal/core/backend/vm"
	"grimlang/internal/core/frontend/ast"
	"grimlang/internal/core/frontend/lexer"
	"grimlang/internal/core/frontend/parser"
	"strconv"
)

var buildins = map[string]byte{
	"add": bytecode.OP_ADD,
	"sub": bytecode.OP_SUB,
	"div": bytecode.OP_DIV,
	"mul": bytecode.OP_MUL,
	"neg": bytecode.OP_NEG,
}

func main() {
	var vm vm.VM
	vm.InitVM()

	code := `
	(add 12 (mul 2 2) 2)
	`
	tokens := lexer.NewLexer(code).Run()
	programm := parser.NewParser(tokens).Run()
	var mainChunk bytecode.Chunk
	mainChunk.InitChunk("main")
	for _, node := range programm.Body {
		Compile(node, &mainChunk)
	}
	mainChunk.DisassemblingChunk()
	vm.RunChunk(&mainChunk)
	vm.PrintTopStack()
	vm.ClearVM()
}

func Compile(node ast.Node, chunk *bytecode.Chunk) {
	switch node := node.(type) {
	case *ast.Number:
		val, err := strconv.ParseFloat(node.Token.Value, 64)
		if err != nil {
			panic(err)
		}
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		enc.Encode(val)
		val_ind := chunk.WriteData(bytecode.Value{Type: object.Float, Object: buf.Bytes()})
		chunk.WriteChunk(bytecode.OP_CONSTANT)
		chunk.WriteChunk(val_ind)
	case *ast.Bool:
		var val bool
		switch node.Token.Value {
		case "true":
			val = true
		case "false":
			val = true
		default:
			panic("unexpected bool value: " + node.Token.Value)
		}
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		enc.Encode(val)
		val_ind := chunk.WriteData(bytecode.Value{Type: object.Bool, Object: buf.Bytes()})
		chunk.WriteChunk(bytecode.OP_CONSTANT)
		chunk.WriteChunk(val_ind)
	case *ast.SymbolExpr:
		op, ok := buildins[node.Symb.Token.Value]
		if !ok {
			panic("Cant find symbol: " + node.Symb.String())
		}
		Compile(node.Args[0], chunk)
		for _, arg := range node.Args[1:] {
			Compile(arg, chunk)
			chunk.WriteChunk(op)
		}
	default:
		panic("Unsupported operation " + node.String())
	}
}
