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
	"add":    bytecode.OP_ADD,
	"sub":    bytecode.OP_SUB,
	"div":    bytecode.OP_DIV,
	"mul":    bytecode.OP_MUL,
	"neg":    bytecode.OP_NEG,
	"not":    bytecode.OP_NOT,
	"lt":     bytecode.OP_LT,
	"gt":     bytecode.OP_GT,
	"leq":    bytecode.OP_LEQ,
	"geq":    bytecode.OP_GEQ,
	"eq":     bytecode.OP_EQ,
	"and":    bytecode.OP_AND,
	"or":     bytecode.OP_OR,
	"concat": bytecode.OP_CONCAT,
	"len":    bytecode.OP_LEN,
	"print":  bytecode.OP_PRINT,
}

func main() {
	var vm vm.VM
	vm.InitVM()

	code := `
	(def str "Hello World")
	(print str)
	(def res (len str))
	(print res)
	`
	tokens := lexer.NewLexer(code).Run()
	programm := parser.NewParser(tokens).Run()
	// fmt.Println(programm.String())
	// println(programm.String())
	var mainChunk bytecode.Chunk
	mainChunk.InitChunk("main")
	for _, node := range programm.Body {
		Compile(node, &mainChunk)
	}
	mainChunk.DisassemblingChunk()
	vm.RunChunk(&mainChunk)
	vm.ClearVM()
}

func Compile(node ast.Node, chunk *bytecode.Chunk) byte {
	switch node := node.(type) {
	case *ast.Number:
		val, err := strconv.ParseFloat(node.Token.Value, 64)
		if err != nil {
			panic(err)
		}
		buf := Encode(val)
		val_ind := chunk.WriteData(bytecode.Value{Type: object.Float, Object: buf.Bytes()})
		chunk.WriteChunk(bytecode.OP_CONSTANT)
		chunk.WriteChunk(val_ind)
		return val_ind
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
		buf := Encode(val)
		val_ind := chunk.WriteData(bytecode.Value{Type: object.Bool, Object: buf.Bytes()})
		chunk.WriteChunk(bytecode.OP_CONSTANT)
		chunk.WriteChunk(val_ind)
		return val_ind
	case *ast.String:
		var val string
		val = node.Token.Value
		buf := Encode(val)
		val_ind := chunk.WriteData(bytecode.Value{Type: object.String, Object: buf.Bytes()})
		chunk.WriteChunk(bytecode.OP_CONSTANT)
		chunk.WriteChunk(val_ind)
		return val_ind
	case *ast.Symbol:
		pointer, ok := chunk.Symbols[node.Token.Value]
		if !ok {
			panic("Symbol " + node.Token.Value + " is undefined")
		}
		chunk.WriteChunk(bytecode.OP_CONSTANT)
		chunk.WriteChunk(pointer)

	case *ast.SymbolExpr:
		op, ok := buildins[node.Symb.Token.Value]
		if !ok {
			panic("Cant find symbol: " + node.Symb.String())

		}
		Compile(node.Args[0], chunk)
		if len(node.Args) == 1 {
			chunk.WriteChunk(op)
		} else {
			for _, arg := range node.Args[1:] {
				Compile(arg, chunk)
				chunk.WriteChunk(op)
			}
		}
	case *ast.DefSF:
		if _, ok := chunk.Symbols[node.Symb.Token.Value]; ok {
			panic(node.Symb.Token.Value + " already defined")
		}
		val_pt := Compile(node.Value, chunk)
		pt := chunk.WriteSymbol(node.Symb.Token.Value, val_pt)
		chunk.WriteChunk(bytecode.OP_DEF)
		chunk.WriteChunk(pt)
	default:
		panic("Unsupported operation " + node.String())
	}
	return 0
}

func Encode[T any](val T) bytes.Buffer {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(val)
	return buf
}
