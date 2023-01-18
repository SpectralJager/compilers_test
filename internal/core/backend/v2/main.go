package main

import (
	"grimlang/internal/core/backend/v2/vm"
)

func main() {
	file := initTestFile()
	v := vm.NewVM(file)
	v.Run()
	v.Disassembly()
	v.DisReg()
}

func initTestFile() []byte {
	var file []byte
	// file header
	header := make([]byte, 64)
	header[0] = 0x45 // E
	header[1] = 0x50 // P
	header[2] = 0x49 // I
	header[3] = 0x45 // E
	header[4] = 0x1  // version 1
	file = append(file, header...)
	// code segment start
	file = append(file, []byte{0xff, 0x0, 0, 0}...)
	// data sigment
	data := make([]byte, 0xff-len(file))
	data[0] = 72
	data[1] = 101
	data[2] = 108
	data[3] = 108
	data[4] = 111
	data[5] = 0
	file = append(file, data...)
	// code sigment
	var code []byte
	code = append(code, vm.InstrToBytes(vm.InstrWithOperOper(vm.LDI, 0, 0))...)
	code = append(code, vm.InstrToBytes(vm.InstrWithOper(vm.PRNT, 0))...)
	code = append(code, vm.InstrToBytes(vm.Instr(vm.HLT))...)
	file = append(file, code...)
	return file
}
