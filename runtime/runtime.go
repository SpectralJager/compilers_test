package runtime

/*
-----------<Enviroment
builtin |>
	int -> import(int)
	main -> import(main)
	exit -> fn[int] <void>
	...
int |>
	add -> fn[...int] <int>
	sub -> fn[...int] <int>
	...
main::builtin|>
	vector -> import(hash1)
	main -> fn[] <void>
hash1::builtin|>
	point -> import(hash2)
	vector -> record<p0::point/point p1::point/point>
	new_vector -> fn[point/point point/point] <vector>
	...
hash2::builtin|>
	point -> record<x0::int x1::int>
	new_point -> fn[int int] <point>
	...
*/
