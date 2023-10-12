package ast

import (
	"fmt"
	"strings"
)

type SymbolAST struct {
	Primary    string     `parser:"@Symbol"`
	Additional *SymbolAST `parser:"('/' @@)?"`
}

func (s *SymbolAST) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s", s.Primary)
	if s.Additional != nil {
		fmt.Fprintf(&buf, "/%s", s.Additional)
	}
	return buf.String()
}

// part of ...
func (s *SymbolAST) atom() {}
func (s *SymbolAST) expr() {}
func (s *SymbolAST) ast()  {}
