package runtime

import "grimlang/backend/asm"

type Frame struct {
	Ip int
	Bp int
	Sp int

	Function asm.Function

	Enviroment Enviroment
}

func NewFrame(fn asm.Function, ip, bp, sp int) Frame {
	return Frame{
		Function:   fn,
		Enviroment: Enviroment{},

		Ip: ip,
		Bp: bp,
		Sp: sp,
	}
}

func (fr *Frame) NextInstruction() *asm.Instruction {
	block, err := fr.Function.Block(fr.Bp)
	if err != nil {
		panic(err)
	}
	instr, err := block.Instruction(fr.Ip)
	if err != nil {
		panic(err)
	}
	fr.Ip++
	return instr
}

func (fr *Frame) SetBlock(index int) {
	_, err := fr.Function.Block(index)
	if err != nil {
		panic(err)
	}
	fr.Bp = index
	fr.Ip = 0
}
