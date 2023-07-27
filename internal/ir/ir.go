package ir

/*
package: *Name*
Constants: // only int, float, string
	0: 12
	1: "some string"
	2: 12.2
	3: 3.14
	...
Globals:
	a int
	b string
	c float
	d bool
	sum func[int int]<int>
	main func[void]<void>
init:
00000000|> global_set a $0;
00000001|> global_set b $1;
00000002|> global_set c $2;
00000003|> global_load b;
00000004|> call_builtin len;
00000005|> global_load a;
00000006|> call_builtin neq;
00000007|> global_save d;
00000008|> call main;
00000009|> exit;
functions:
=== main:
Arguments:
	void
Locals:
	result int
Body:
00000000|> load $12;
00000001|> load $11;
00000002|> call sum;
00000003|> local_save result;
00000004|> local_load result;
00000005|> call_builtin print;
00000006|> ret;
=== main:
Arguments:
	arg1 int
	arg2 int
Locals:
	void
Body:
00000000|> call_builtin iadd;
00000001|> ret;
*/
