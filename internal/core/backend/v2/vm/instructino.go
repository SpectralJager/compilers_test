package vm

import "encoding/binary"

type Instruction uint32

func Instr(op uint8) Instruction {
	temp := uint32(op)
	return Instruction(temp)
}

func InstrWithOper(op uint8, operand1 uint32) Instruction {
	temp := uint32(op) | (operand1 << 8)
	return Instruction(temp)
}

func InstrWithOperOper(op uint8, operand1 uint8, operand2 uint16) Instruction {
	temp := uint32(0) | uint32(op) | (uint32(operand1) << 8) | (uint32(operand2) << 16)
	return Instruction(temp)
}

func InstrWithOperOperOper(op uint8, operand1 uint8, operand2 uint8, operand3 uint8) Instruction {
	temp := uint32(0) | uint32(op) | (uint32(operand1) << 8) | (uint32(operand2) << 16) | (uint32(operand3) << 24)
	return Instruction(temp)
}

func (i *Instruction) GetOp() uint8 {
	return uint8(uint32(*i) & uint32(0xff))
}

func (i *Instruction) GetOper24() uint32 {
	return (uint32(*i) >> 8) & uint32(0xffffff)
}

func (i *Instruction) GetFirstOper8() uint8 {
	return uint8((uint32(*i) >> 8) & uint32(0xff))
}

func (i *Instruction) GetSecondOper16() uint16 {
	return uint16((uint32(*i) >> 16) & uint32(0xffff))
}

func (i *Instruction) GetSecondOper8() uint8 {
	return uint8((uint32(*i) >> 16) & uint32(0xff))
}

func (i *Instruction) GetThirdOper8() uint8 {
	return uint8((uint32(*i) >> 24) & uint32(0xff))
}

func InstrToBytes(instr Instruction) []byte {
	a := make([]byte, 4)
	binary.BigEndian.PutUint32(a, uint32(instr))
	a[0], a[1], a[2], a[3] = a[3], a[2], a[1], a[0]
	return a
}
