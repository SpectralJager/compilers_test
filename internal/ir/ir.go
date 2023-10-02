package ir

/*
types |>
	int -> INT
	bool -> BOOL
init |>
	var.new n int
	const.push 40
	var.pop n
	call main
functions |>
fib:
	var.new n int
	var.pop n
	var.push n
	const.push 2
	reg.set ar 2
	call int/lt
	jmp.false if_else_0 ; if !grim.pop().(*bool) {goto if_else_0}
	block.start
	var.push n;
	reg.set rt 1
	ret
	block.end
	:if_else_0
	block.start
	var.push n
	const.push 1
	reg.set ar 2
	call int/sub
	var.push n
	const.push 2
	reg.set ar 2
	call int/sub
	call int/add
	call fib
	reg.set rt 1
	ret
	block.end
	:if_end_0
*/
