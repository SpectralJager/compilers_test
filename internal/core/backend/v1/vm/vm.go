package vm

import (
	"fmt"
	"grimlang/internal/core/backend/v1/bytecode"
	"grimlang/internal/core/backend/v1/chunk"
)

var buildins = map[string]func(args []any) any{
	"println": func(args []any) any { fmt.Println(args...); return nil },
	"add": func(args []any) any {
		res := 0.0
		for i := len(args) - 1; i >= 0; i-- {
			res += args[i].(float64)
		}
		return res
	},
	"sub": func(args []any) any {
		res := args[len(args)-1].(float64)
		for i := len(args) - 2; i >= 0; i-- {
			res -= args[i].(float64)
		}
		return res
	},
	"mul": func(args []any) any {
		res := args[len(args)-1].(float64)
		for i := len(args) - 2; i >= 0; i-- {
			res *= args[i].(float64)
		}
		return res
	},
	"div": func(args []any) any {
		res := args[len(args)-1].(float64)
		for i := len(args) - 2; i >= 0; i-- {
			res /= args[i].(float64)
		}
		return res
	},
	"pow": func(args []any) any {
		if len(args) != 2 {
			panic("func pow expected 2 args of number")
		}
		val := args[1].(float64)
		res := 1.0
		n := args[0].(float64)
		for i := 0; i < int(n); i++ {
			res *= val
		}
		return res
	},
	"lt": func(args []any) any {
		val := args[len(args)-1].(float64)
		for i := len(args) - 2; i >= 0; i-- {
			if val < args[i].(float64) {
				val = args[i].(float64)
				continue
			}
			return false
		}
		return true
	},
	"gt": func(args []any) any {
		val := args[len(args)-1].(float64)
		for i := len(args) - 2; i >= 0; i-- {
			if val > args[i].(float64) {
				val = args[i].(float64)
				continue
			}
			return false
		}
		return true
	},
	"leq": func(args []any) any {
		val := args[len(args)-1].(float64)
		for i := len(args) - 2; i >= 0; i-- {
			if val <= args[i].(float64) {
				val = args[i].(float64)
				continue
			}
			return false
		}
		return true
	},
	"geq": func(args []any) any {
		val := args[len(args)-1].(float64)
		for i := len(args) - 2; i >= 0; i-- {
			if val >= args[i].(float64) {
				val = args[i].(float64)
				continue
			}
			return false
		}
		return true
	},
	"eq": func(args []any) any {
		val := args[len(args)-1].(float64)
		for i := len(args) - 2; i >= 0; i-- {
			if val == args[i].(float64) {
				val = args[i].(float64)
				continue
			}
			return false
		}
		return true
	},
}

type VM struct {
	Stack []any
	SP    int
	Env   map[string]any
}

func NewVM() *VM {
	return &VM{
		Stack: make([]any, 65535),
		SP:    0,
		Env:   make(map[string]any),
	}
}

func (vm *VM) ExecuteChunk(ch chunk.Chunk) {
	for _, bt := range ch.Bytecodes {
		switch bt.GetOpcode() {
		case bytecode.OP_HLT:
			return
		case bytecode.OP_LOAD_CONST:
			vm.Push(bt.GetwArgs())
		case bytecode.OP_CALL:
			args := bt.GetwArgs()
			nargs, ok := args.(map[string]any)["nargs"].(int)
			if !ok {
				panic("for op_call expected nargs arg key")
			}
			op, ok := args.(map[string]any)["symbol"].(string)
			if !ok {
				panic("for op_call expected symbol arg key")
			}
			if _, ok := buildins[op]; ok {
				_args := []any{}
				for i := 0; i < nargs; i++ {
					_args = append(_args, vm.Pop())
				}
				val := buildins[op](_args)
				if val != nil {
					vm.Push(val)
				}
			} else if ch, ok := vm.Env[op]; ok {
				var copyStack []any
				copy(copyStack, vm.Stack)
				copyEnv := make(map[string]any)
				for k, v := range vm.Env {
					copyEnv[k] = v
				}
				vm.ExecuteChunk(*ch.(*chunk.Chunk))
				copy(vm.Stack, copyStack)
				val, ok := vm.Env["ret"]
				if ok {
					vm.Push(val)
				}
				vm.Env = make(map[string]any)
				for k, v := range copyEnv {
					vm.Env[k] = v
				}
			} else {
				panic(fmt.Sprintf("symbol %s is undefined!", op))
			}
		case bytecode.OP_IF:
			logVal := vm.Pop().(bool)
			args := bt.GetwArgs().(map[string]any)
			if logVal {
				ch := args["true"]
				var copyStack []any
				copy(copyStack, vm.Stack)
				copyEnv := make(map[string]any)
				for k, v := range vm.Env {
					copyEnv[k] = v
				}
				vm.ExecuteChunk(*ch.(*chunk.Chunk))
				copy(vm.Stack, copyStack)
				for k, v := range copyEnv {
					val := vm.Env[k]
					if val != v {
						copyEnv[k] = val
						continue
					}
				}
				vm.Env = make(map[string]any)
				for k, v := range copyEnv {
					vm.Env[k] = v
				}
			} else {
				ch := args["false"]
				var copyStack []any
				copy(copyStack, vm.Stack)
				copyEnv := make(map[string]any)
				for k, v := range vm.Env {
					copyEnv[k] = v
				}
				vm.ExecuteChunk(*ch.(*chunk.Chunk))
				copy(vm.Stack, copyStack)
				for k, v := range copyEnv {
					val := vm.Env[k]
					if val != v {
						copyEnv[k] = val
						continue
					}
				}
				vm.Env = make(map[string]any)
				for k, v := range copyEnv {
					vm.Env[k] = v
				}
			}
		case bytecode.OP_SAVE_NAME:
			symb := bt.GetwArgs().(string)
			_, ok := buildins[symb]
			if ok {
				panic(fmt.Sprintf("symbol %s is buildin function! use another name", symb))
			}
			if ch.Name == "main" {
				_, ok = vm.Env[symb]
				if ok {
					panic(fmt.Sprintf("symbol %s already defined! use another name", symb))
				}
			}
			val := vm.Pop()
			vm.Env[symb] = val
		case bytecode.OP_SET_NAME:
			symb := bt.GetwArgs().(string)
			_, ok := buildins[symb]
			if ok {
				panic(fmt.Sprintf("symbol %s is buildin function! use another name", symb))
			}
			_, ok = vm.Env[symb]
			if !ok {
				panic(fmt.Sprintf("symbol %s undefined!", symb))
			}
			val := vm.Pop()
			vm.Env[symb] = val
		case bytecode.OP_LOAD_NAME:
			symb := bt.GetwArgs().(string)
			_, ok := buildins[symb]
			if ok {
				panic(fmt.Sprintf("cant use symbol %s as argument!", symb))
			}
			val, ok := vm.Env[symb]
			if !ok {
				panic(fmt.Sprintf("symbol %s undefined!", symb))
			}
			vm.Push(val)
		case bytecode.OP_SAVE_FN:
			args := bt.GetwArgs().(map[string]any)
			symb := args["symbol"].(string)
			_, ok := buildins[symb]
			if ok {
				panic(fmt.Sprintf("symbol %s is buildin function! use another name", symb))
			}
			_, ok = vm.Env[symb]
			if ok {
				panic(fmt.Sprintf("symbol %s already defined! use another name", symb))
			}
			body, ok := args["body"]
			if !ok {
				panic(fmt.Sprintf("body of %s undefined!", symb))
			}
			vm.Env[symb] = body
		case bytecode.OP_RET:
			args := bt.GetwArgs().(bool)
			if args {
				val := vm.Pop()
				vm.Env["ret"] = val
			}
			return
		default:
			panic("unsupported bytecode's operation: " + bt.String())
		}
	}
}

func (vm *VM) Push(val any) {
	// fmt.Printf("%v-%T %d\n", val, val, vm.SP)
	vm.SP += 1
	vm.Stack[vm.SP-1] = val
}

func (vm *VM) Pop() any {
	if vm.SP == 0 {
		panic("trying to pop empty stack")
	}
	val := vm.Stack[vm.SP-1]
	vm.SP -= 1
	return val
}

func (vm *VM) TraiceStack() {
	fmt.Printf("------VM Stack------\n")
	for i := vm.SP; i > 0; i-- {
		fmt.Printf("#%d\t%v\n", i-1, vm.Stack[i-1])
	}
	fmt.Printf("--------------------\n\n")
}
