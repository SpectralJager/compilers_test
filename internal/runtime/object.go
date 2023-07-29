package runtime

import "fmt"

type Object interface {
	fmt.Stringer
	object()
}

func (*Integer) object() {}
func (*Float) object()   {}
func (*String) object()  {}

type Integer struct {
	Value int
}

type Float struct {
	Value float64
}

type String struct {
	Value string
}

type Varible struct {
	Name  string
	Value Object
}

// Stringers
func (s *Float) String() string {
	return fmt.Sprintf("float:%f", s.Value)
}
func (s *Integer) String() string {
	return fmt.Sprintf("int:%d", s.Value)
}
func (s *String) String() string {
	return fmt.Sprintf("string:%s", s.Value)
}
func (s *Varible) String() string {
	return fmt.Sprintf("varible:%s=%s", s.Name, s.Value)
}
