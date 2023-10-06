package ir

/*
module:

	name -> ...
	path -> ...
	imports |>
		"..." as ...

global |>

	block.begin global
	var.new a int
	var.set a 12
	call main 0
	block.end

main |>

	block.begin fn
	var.new res int
	const.push 40
	call fib 1
	var.pop res
	var.push res
	callb printf 1
	block.end

fib |>

	block.begin fn
	var.new n int
	var.pop n
	var.push n
	const.push 2
	callb int/lt 2
	block.begin local
	var.push n
	stack.type int
	ret 1
	block.end
	var.push n
	const.push 1
	callb int/sub 2
	var.push n
	const.push 2
	callb int/sub 2
	callb int/add 2
	call fib 1
	stack.type int
	ret 1
	block.end
*/
