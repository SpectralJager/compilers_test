package ir

import (
	"fmt"
	"log"
	"strings"
)

type Module struct {
	Name   string
	Path   string
	Chunks []*Chunk
}

func NewModule(name string, path string) *Module {
	return &Module{
		Name: name,
		Path: path,
	}
}

func (m *Module) AddChunk(ch *Chunk) *Module {
	if ch := m.GetChunk(ch.Name); ch != nil {
		log.Fatalf("chunk %v already exists", ch.Name)
	}
	m.Chunks = append(m.Chunks, ch)
	return m
}

func (m *Module) GetChunk(name string) *Chunk {
	for _, chunk := range m.Chunks {
		if chunk.Name == name {
			return chunk
		}
	}
	return nil
}

func (m *Module) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "=== Module '%s'\n", m.Name)
	fmt.Fprintf(&buf, "Path '%s'\n", m.Path)
	fmt.Fprint(&buf, "Chunks |>\n")
	if len(m.Chunks) == 0 {
		fmt.Fprint(&buf, "\tempty\n")
	} else {
		for _, chunk := range m.Chunks {
			fmt.Fprintf(&buf, "%s", chunk.String())
		}
	}
	fmt.Fprintf(&buf, "=== end '%s'\n", m.Name)
	return buf.String()
}
