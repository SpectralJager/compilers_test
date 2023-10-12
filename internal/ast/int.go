package ast

import "fmt"

type IntAST struct {
	Value int `parser:"@Integer"`
}

func (a *IntAST) String() string {
	return fmt.Sprintf("%d", a.Value)
}

// part of ...
func (a *IntAST) expr() {}
func (a *IntAST) atom() {}
func (a *IntAST) ast()  {}
