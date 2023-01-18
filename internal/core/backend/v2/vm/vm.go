package vm

import (
	"fmt"
)

type VM struct {
	regisers  [32]uint32
	heap      []byte
	pc        uint32
	pd        uint32
	code      []byte
	remainder uint32
	eq_flag   bool
}

func NewVM(code []byte) *VM {
	vm := VM{
		code:    code,
		eq_flag: false,
	}

	var header string
	for i := 0; i < 4; i++ {
		header += string(vm.nextByte())
	}
	if "EPIE" != header {
		panic("wrong file header")
	}
	vm.pc = 64
	code_start := vm.next4Bytes()
	vm.pd = vm.pc
	vm.pc = code_start

	return &vm
}

func (vm *VM) Run() {
	var isHlt bool = false
	for !isHlt && vm.pc < uint32(len(vm.code)) {
		opcode := vm.nextByte()
		switch opcode {
		case HLT:
			isHlt = true
		case LDI:
			reg := vm.nextByte()
			val := vm.next2Bytes()
			vm.regisers[reg] = uint32(val)
		case ALOC:
			bytes := vm.regisers[vm.next3Bytes()]
			for i := uint32(0); i < bytes; i++ {
				vm.heap = append(vm.heap, 0)
			}
		case PRNT:
			st := vm.regisers[vm.next3Bytes()]
			temp := ""
			for i := vm.pd + st; vm.code[i] != 0x00; i++ {
				temp += string(vm.code[i])
			}
			fmt.Print(temp)
		case ADD:
			reg1 := vm.nextByte()
			reg2 := vm.nextByte()
			reg3 := vm.nextByte()
			vm.regisers[reg3] = vm.regisers[reg1] + vm.regisers[reg2]
		case SUB:
			reg1 := vm.nextByte()
			reg2 := vm.nextByte()
			reg3 := vm.nextByte()
			vm.regisers[reg3] = vm.regisers[reg1] - vm.regisers[reg2]
		case MUL:
			reg1 := vm.nextByte()
			reg2 := vm.nextByte()
			reg3 := vm.nextByte()
			vm.regisers[reg3] = vm.regisers[reg1] * vm.regisers[reg2]
		case DIV:
			reg1 := vm.nextByte()
			reg2 := vm.nextByte()
			reg3 := vm.nextByte()
			vm.regisers[reg3] = vm.regisers[reg1] / vm.regisers[reg2]
			vm.remainder = vm.regisers[reg1] % vm.regisers[reg2]
		case EQ:
			vm.eq_flag = false
			reg1 := vm.regisers[vm.nextByte()]
			reg2 := vm.regisers[vm.nextByte()]
			if reg1 == reg2 {
				vm.eq_flag = true
			}
		case JMP:
			vm.pc = vm.regisers[vm.nextByte()]
			continue
		case JMF:
			vm.pc += vm.next3Bytes()
			continue
		case JMB:
			vm.pc -= vm.next3Bytes()
			continue
		case JEQ:
			if vm.eq_flag {
				vm.pc = vm.regisers[vm.nextByte()]
				continue
			}
		case NOP:
			continue
		default:
			panic(fmt.Sprintf("Unexpected opcode: %02x", opcode))
		}
	}
}

func (vm *VM) nextByte() byte {
	temp := vm.code[vm.pc]
	vm.pc += 1
	return temp
}

func (vm *VM) next2Bytes() uint16 {
	temp := vm.code[vm.pc : vm.pc+2]
	vm.pc += 2
	return uint16(temp[0]) | (uint16(temp[1]) << 8)
}
func (vm *VM) next3Bytes() uint32 {
	temp := vm.code[vm.pc : vm.pc+3]
	vm.pc += 3
	return uint32(temp[0]) | uint32(temp[1])<<8 | uint32(temp[2])<<16
}

func (vm *VM) next4Bytes() uint32 {
	temp := vm.code[vm.pc : vm.pc+4]
	vm.pc += 4
	return uint32(temp[0]) | uint32(temp[1])<<8 | uint32(temp[2])<<16 | uint32(temp[3])<<24
}

func (vm *VM) next8Bytes() uint64 {
	temp := vm.code[vm.pc : vm.pc+8]
	vm.pc += 8
	return uint64(temp[0]) | uint64(temp[1])<<8 | uint64(temp[2])<<16 | uint64(temp[3])<<24 | uint64(temp[4])<<32 | uint64(temp[5])<<40 | uint64(temp[6])<<48 | uint64(temp[7])<<56
}

func (vm *VM) Disassembly() {
	temp := vm.pc
	vm.pc = 64
	code_seg := vm.next4Bytes()
	vm.pc = temp

	fmt.Print("\nData\n")
	for i, v := range vm.code[68:code_seg] {
		fmt.Printf("%08x: %02x\n", 68+i, v)
	}

	fmt.Print("\nCode\n")
	for i, v := range vm.code[code_seg:] {
		fmt.Printf("%08x: %02x\n", int(code_seg)+i, v)
	}
}

func (vm *VM) DisReg() {
	fmt.Print("\nRegisters\n")
	for i, v := range vm.regisers {
		fmt.Printf("%02d: %08x\n", i, v)
	}
	fmt.Printf("remainder: %08x\n", vm.remainder)

	fmt.Print("\nFlags\n")
	fmt.Printf("eq: %v\n", vm.eq_flag)
}

func (vm *VM) DisHeap() {
	fmt.Print("\nHeap\n")
	for i, v := range vm.heap {
		fmt.Printf("%08x: %02x\n", i, v)
	}
}
