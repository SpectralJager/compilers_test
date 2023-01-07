package main

import (
	"fmt"
	"grimlang/internal/core/backend/bytecode"
	"grimlang/internal/core/backend/chunk"
	"grimlang/internal/core/backend/vm"
	"grimlang/internal/core/frontend/ast"
	"grimlang/internal/core/frontend/lexer"
	"grimlang/internal/core/frontend/parser"
	"strconv"
)

func main() {
	code := `
	(fn recur [c] (
		(if (leq c 0) (
			(ret)
		))
		(println c)
		(recur (sub c 1))

	))
	(def cnt 10)
	(recur cnt)
	`
	programm := parser.NewParser(
		lexer.NewLexer(code).Run(),
	).Run()
	mainChunk := chunk.NewChunk("main", 0x0000)
	for _, node := range programm.Body {
		Compile(node, mainChunk)
	}
	hltBt := bytecode.NewBytecode(bytecode.OP_HLT, nil)
	mainChunk.WriteBytecode(*hltBt)
	mainChunk.Disassembly()
	vir := vm.NewVM()
	vir.ExecuteChunk(*mainChunk)
	vir.TraiceStack()
}

func Compile(node ast.Node, c *chunk.Chunk) {
	switch node := node.(type) {
	case *ast.Number:
		val, err := strconv.ParseFloat(node.Token.Value, 64)
		if err != nil {
			panic(err)
		}
		bt := bytecode.NewBytecode(bytecode.OP_LOAD_CONST, val)
		c.WriteBytecode(*bt)
	case *ast.Bool:
		var val bool
		switch node.Token.Value {
		case "true":
			val = true
		case "false":
			val = false
		}
		bt := bytecode.NewBytecode(bytecode.OP_LOAD_CONST, val)
		c.WriteBytecode(*bt)
	case *ast.Symbol:
		bt := bytecode.NewBytecode(bytecode.OP_LOAD_NAME, node.Token.Value)
		c.WriteBytecode(*bt)
	case *ast.SymbolExpr:
		for _, arg := range node.Args {
			Compile(arg, c)
		}
		bt := bytecode.NewBytecode(bytecode.OP_CALL, map[string]any{"symbol": node.Symb.Token.Value, "nargs": len(node.Args)})
		c.WriteBytecode(*bt)
	case *ast.DefSF:
		Compile(node.Value, c)
		bt := bytecode.NewBytecode(bytecode.OP_SAVE_NAME, node.Symb.Token.Value)
		c.WriteBytecode(*bt)
	case *ast.SetSF:
		Compile(node.Value, c)
		bt := bytecode.NewBytecode(bytecode.OP_SET_NAME, node.Symb.Token.Value)
		c.WriteBytecode(*bt)
	case *ast.RetSF:
		bt := bytecode.NewBytecode(bytecode.OP_RET, false)
		if node.Value != nil {
			Compile(node.Value, c)
			bt = bytecode.NewBytecode(bytecode.OP_RET, true)
		}
		c.WriteBytecode(*bt)
	case *ast.FnSF:
		symb := node.Symb.Token.Value
		symbChunk := chunk.NewChunk(symb, 0x0000)
		for i := len(node.Args) - 1; i >= 0; i-- {
			arg := node.Args[i]
			bt := bytecode.NewBytecode(bytecode.OP_SAVE_NAME, arg.Token.Value)
			symbChunk.WriteBytecode(*bt)
		}
		for _, nd := range node.Body {
			Compile(nd, symbChunk)
		}
		symbChunk.Disassembly()
		bt := bytecode.NewBytecode(bytecode.OP_SAVE_FN, map[string]any{"symbol": node.Symb.Token.Value, "body": symbChunk})
		c.WriteBytecode(*bt)
	case *ast.IfSF:
		Compile(node.Log, c)
		TrueChunk := chunk.NewChunk("trueChunk", 0x0000)
		for _, nd := range node.TrueValue {
			Compile(nd, TrueChunk)
		}
		// TrueChunk.Disassembly()
		FalseChunk := chunk.NewChunk("falseChunk", 0x0000)
		for _, nd := range node.FalseValue {
			Compile(nd, FalseChunk)
		}
		// FalseChunk.Disassembly()
		bt := bytecode.NewBytecode(bytecode.OP_IF, map[string]any{"true": TrueChunk, "false": FalseChunk})
		c.WriteBytecode(*bt)
	default:
		panic("unexpected node type" + fmt.Sprintf("%T", node))
	}
}
