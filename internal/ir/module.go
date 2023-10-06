package ir

import (
	"fmt"
	"strings"
)

type Module struct {
	Name    string
	Path    string
	Imports map[string]string
	Chunks  []Chunk
}

func NewModule(name string, path string, imports map[string]string, chunks ...Chunk) *Module {
	return &Module{
		Name:    name,
		Path:    path,
		Imports: imports,
		Chunks:  chunks,
	}
}

func (md *Module) String() string {
	var buf strings.Builder
	fmt.Fprint(&buf, "module: \n")
	fmt.Fprintf(&buf, "\tname -> %s\n", md.Name)
	fmt.Fprintf(&buf, "\tpath -> %s\n", md.Path)
	for _, ch := range md.Chunks {
		fmt.Fprintf(&buf, "%s\n", ch.String())
	}
	return buf.String()
}

func (md *Module) AddChunks(chunks ...Chunk) {
	md.Chunks = append(md.Chunks, chunks...)
}
