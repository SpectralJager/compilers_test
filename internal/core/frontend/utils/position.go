package utils

import "fmt"

type Position struct {
	Line   int
	Column int
}

func (p *Position) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}
