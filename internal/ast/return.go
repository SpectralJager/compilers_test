package ast

import "fmt"

type ReturnAST struct {
	Value EXPR `parser:"'@ret' @@ ';'"`
}

func (r *ReturnAST) String() string {
	return fmt.Sprintf("@ret %s", r.Value)
}

// part of ...
func (r *ReturnAST) locl() {}
func (r *ReturnAST) ast()  {}
