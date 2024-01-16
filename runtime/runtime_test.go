package runtime

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	intType := NewIntType()

	fmt.Println(intType.Name())
	fmt.Println(intType.Compare(NewIntType()))
	fmt.Println(intType.Subtype())
}
