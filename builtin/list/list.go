package builtin_list

import "grimlang/runtime"

func NewBuiltinListEnv() runtime.Enviroment {
	env := runtime.NewEnviroment("list", nil)
	env.Insert(
		runtime.NewBuiltin(
			"len",
			runtime.NewFunctionType(
				runtime.NewIntType(),
				runtime.NewListType(runtime.NewAnyType()),
			),
			Len,
		),
	)
	env.Insert(
		runtime.NewBuiltin(
			"get",
			runtime.NewFunctionType(
				runtime.NewAnyType(),
				runtime.NewListType(runtime.NewAnyType()),
				runtime.NewIntType(),
			),
			Get,
		),
	)
	env.Insert(
		runtime.NewBuiltin(
			"insert",
			runtime.NewFunctionType(
				runtime.NewListType(runtime.NewAnyType()),
				runtime.NewListType(runtime.NewAnyType()),
				runtime.NewIntType(),
				runtime.NewAnyType(),
			),
			Insert,
		),
	)
	env.Insert(
		runtime.NewBuiltin(
			"append",
			runtime.NewFunctionType(
				runtime.NewListType(runtime.NewAnyType()),
				runtime.NewListType(runtime.NewAnyType()),
				runtime.NewAnyType(),
			),
			Append,
		),
	)
	return env
}

func Len(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) != 1 {
		panic("list/len: expect 1 input")
	}
	return runtime.NewIntLit(
		int64(inputs[0].Len()),
	)
}

func Get(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) != 2 {
		panic("list/get: expect 2 inputs")
	}
	return inputs[0].Item(
		int(inputs[1].ValueInt()),
	)
}

func Insert(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) != 3 {
		panic("list/insert: expect 3 inputs")
	}
	list := inputs[0]
	index := inputs[1]
	value := inputs[2]
	items := []runtime.Litteral{}
	for i := 0; i < list.Len(); i++ {
		items = append(items, list.Item(i))
	}
	switch {
	case list.Len() <= int(index.ValueInt()):
		items = append(items, value)
	case list.Len() > int(index.ValueInt()):
		items = append(
			items[:index.ValueInt()],
			append([]runtime.Litteral{value}, items[index.ValueInt():]...)...,
		)
	}
	return runtime.NewListLit(
		list.Type().Item(),
		items...,
	)
}

func Append(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) != 2 {
		panic("list/append: expect 2 inputs")
	}
	list := inputs[0]
	value := inputs[1]
	items := []runtime.Litteral{}
	for i := 0; i < list.Len(); i++ {
		items = append(items, list.Item(i))
	}
	items = append(items, value)
	return runtime.NewListLit(
		list.Type().Item(),
		items...,
	)
}
