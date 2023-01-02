package vm

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"grimlang/internal/core/backend/bytecode"
	"grimlang/internal/core/backend/object"
)

type VM struct {
	Chunk *bytecode.Chunk
	Stack []bytecode.Value
}

func (vm *VM) InitVM() {
	vm.ClearVM()
}

func (vm *VM) ClearVM() {
	vm.Stack = make([]bytecode.Value, 0)
}

func (vm *VM) RunChunk(chunk *bytecode.Chunk) error {
	vm.Chunk = chunk
	for ip := 0; ip < len(vm.Chunk.Code); ip++ {
		var buf bytes.Buffer
		dec := gob.NewDecoder(&buf)
		enc := gob.NewEncoder(&buf)
		switch vm.Chunk.Code[ip] {
		case bytecode.OP_RETURN:
			return nil
		case bytecode.OP_CONSTANT:
			ip += 1
			val_pt := vm.Chunk.Code[ip]
			val := vm.Chunk.Data[val_pt]
			vm.Push(val)
		case bytecode.OP_NEG:
			val := vm.Pop()
			_, err := buf.Write(val.Object)
			if err != nil {
				panic(err)
			}
			var el float64
			err = dec.Decode(&el)
			if err != nil {
				panic(err)
			}
			el = -el
			err = enc.Encode(&el)
			if err != nil {
				panic(err)
			}
			val.Object = buf.Bytes()
			vm.Push(val)

		case bytecode.OP_ADD:
			a := vm.Pop()
			b := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b float64
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			el_a += el_b
			buf.Reset()
			err = enc.Encode(&el_a)
			if err != nil {
				panic(err)
			}
			a.Object = buf.Bytes()
			vm.Push(a)
		case bytecode.OP_SUB:
			a := vm.Pop()
			b := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b float64
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			el_a -= el_b
			buf.Reset()
			err = enc.Encode(&el_a)
			if err != nil {
				panic(err)
			}
			a.Object = buf.Bytes()
			vm.Push(a)
		case bytecode.OP_MUL:
			a := vm.Pop()
			b := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b float64
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			el_a *= el_b
			buf.Reset()
			err = enc.Encode(&el_a)
			if err != nil {
				panic(err)
			}
			a.Object = buf.Bytes()
			vm.Push(a)
		case bytecode.OP_DIV:
			a := vm.Pop()
			b := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b float64
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			el_a /= el_b
			buf.Reset()
			err = enc.Encode(&el_a)
			if err != nil {
				panic(err)
			}
			a.Object = buf.Bytes()
			vm.Push(a)
		case bytecode.OP_NOT:
			a := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			var el_a bool
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			el_a = !el_a
			err = enc.Encode(&el_a)
			if err != nil {
				panic(err)
			}
			a.Object = buf.Bytes()
			vm.Push(a)
		case bytecode.OP_LT:
			b := vm.Pop()
			a := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b float64
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			temp := el_a < el_b
			err = enc.Encode(&temp)
			if err != nil {
				panic(err)
			}
			vm.Push(bytecode.Value{Type: object.Bool, Object: buf.Bytes()})
		case bytecode.OP_GT:
			b := vm.Pop()
			a := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b float64
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			temp := el_a > el_b
			err = enc.Encode(&temp)
			if err != nil {
				panic(err)
			}
			vm.Push(bytecode.Value{Type: object.Bool, Object: buf.Bytes()})
		case bytecode.OP_LEQ:
			b := vm.Pop()
			a := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b float64
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			temp := el_a <= el_b
			err = enc.Encode(&temp)
			if err != nil {
				panic(err)
			}
			vm.Push(bytecode.Value{Type: object.Bool, Object: buf.Bytes()})
		case bytecode.OP_GEQ:
			b := vm.Pop()
			a := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b float64
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			temp := el_a >= el_b
			err = enc.Encode(&temp)
			if err != nil {
				panic(err)
			}
			vm.Push(bytecode.Value{Type: object.Bool, Object: buf.Bytes()})
		case bytecode.OP_EQ:
			b := vm.Pop()
			a := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b float64
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			temp := el_a == el_b
			err = enc.Encode(&temp)
			if err != nil {
				panic(err)
			}
			vm.Push(bytecode.Value{Type: object.Bool, Object: buf.Bytes()})
		case bytecode.OP_AND:
			b := vm.Pop()
			a := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b bool
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			temp := el_a && el_b
			err = enc.Encode(&temp)
			if err != nil {
				panic(err)
			}
			vm.Push(bytecode.Value{Type: object.Bool, Object: buf.Bytes()})
		case bytecode.OP_OR:
			b := vm.Pop()
			a := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b bool
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			temp := el_a || el_b
			err = enc.Encode(&temp)
			if err != nil {
				panic(err)
			}
			vm.Push(bytecode.Value{Type: object.Bool, Object: buf.Bytes()})
		case bytecode.OP_CONCAT:
			b := vm.Pop()
			a := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			_, err = buf.Write(b.Object)
			if err != nil {
				panic(err)
			}
			var el_a, el_b string
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			err = dec.Decode(&el_b)
			if err != nil {
				panic(err)
			}
			el_a += el_b
			err = enc.Encode(&el_a)
			if err != nil {
				panic(err)
			}
			a.Object = buf.Bytes()
			vm.Push(a)
		case bytecode.OP_LEN:
			a := vm.Pop()
			_, err := buf.Write(a.Object)
			if err != nil {
				panic(err)
			}
			var el_a string
			err = dec.Decode(&el_a)
			if err != nil {
				panic(err)
			}
			temp := float64(len(el_a))
			err = enc.Encode(&temp)
			if err != nil {
				panic(err)
			}
			vm.Push(bytecode.Value{Type: object.Float, Object: buf.Bytes()})
		}
	}
	return nil
}

func (vm *VM) Push(val bytecode.Value) {
	vm.Stack = append(vm.Stack, val)
}

func (vm *VM) Pop() bytecode.Value {
	if len(vm.Stack) != 0 {
		val := vm.Stack[len(vm.Stack)-1]
		vm.Stack = vm.Stack[:len(vm.Stack)-1]
		return val
	}
	return bytecode.Value{}
}

func (vm *VM) PrintTopStack() {
	var buf bytes.Buffer
	temp := vm.Stack[len(vm.Stack)-1]
	_, err := buf.Write(temp.Object)
	if err != nil {
		panic(err)
	}
	switch temp.Type {
	case object.Float:
		decode[float64](&buf)
	case object.Bool:
		decode[bool](&buf)
	case object.String:
		decode[string](&buf)
	default:
		panic("unsupported value type: " + fmt.Sprint(temp.Type))
	}
}

func decode[T any](buf *bytes.Buffer) {
	var val T
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&val)
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
