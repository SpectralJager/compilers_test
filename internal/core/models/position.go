package models

import "fmt"

type Position struct {
	line   int
	column int
}

func (p *Position) String() string {
	return fmt.Sprintf("%d:%d", p.line, p.column)
}
