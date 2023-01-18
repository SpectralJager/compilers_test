package vm

const (
	NOP uint8 = iota
	HLT
	IGL

	LDI // ldi %r #v
	LDX // ldx %r %r

	ADD // add %r1 %r2 %r_res
	SUB // sub %r1 %r2 %r_res
	MUL // mul %r1 %r2 %r_res
	DIV // div %r1 %r2 %r_res $remainder

	JMP // jmp %r
	JMF // jMF #v ; jump forward
	JMB // jMB #v ; jump backward
	JEQ //

	EQ

	ALOC // aloc %r

	PRNT // prnt %r
)
