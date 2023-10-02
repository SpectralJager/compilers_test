package ast

type IfAST struct {
	IfCondition EXPR    `parser:"'@if' @@ "`
	IfBody      []LOCAL `parser:"'{' @@+ "`
	ElseBody    []LOCAL `parser:"('else' '=>' @@+)? '}'"`
}

func (f *IfAST) String() string {
	return ""
}

// part of ...
func (f *IfAST) locl() {}
