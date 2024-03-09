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
		Enviroment: NewEnviroment(fn.Vars),

		Ip: ip,
		Bp: bp,
		Sp: sp,
	}
}

func (fr *Frame) NextInstruction() (*asm.Instruction, error) {
	block, err := fr.Function.Block(fr.Bp)
	if err != nil {
		return &asm.Instruction{}, err
	}
	instr, err := block.Instruction(fr.Ip)
	if err != nil {
		return &asm.Instruction{}, err
	}
	fr.Ip++
	return instr, nil
}

func (fr *Frame) SetBlock(index int) error {
	_, err := fr.Function.Block(index)
	if err != nil {
		return err
	}
	fr.Bp = index
	fr.Ip = 0
	return nil
}
