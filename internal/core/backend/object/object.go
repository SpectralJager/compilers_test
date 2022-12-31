package object

type ObjectType int

const (
	Float ObjectType = iota
	Bool
)

type Object struct {
	Bytes []byte
}
