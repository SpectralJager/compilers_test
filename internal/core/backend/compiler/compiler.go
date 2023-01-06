package main

import (
	"grimlang/internal/core/backend/chunk"
	"grimlang/internal/core/frontend/ast"
)

func main() {
}

func Compile(node ast.Node, ch chunk.Chunk) {
	switch node := node.(type) {
	case *ast.Number:

	}
}
