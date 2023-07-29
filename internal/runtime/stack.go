package runtime

/*
Stack:
0 -> int:40
*/

type Stack []Object

func (s *Stack) Len() int { return len(*s) }
func (s *Stack) Push(obj Object) {
	*s = append(*s, obj)
}
func (s *Stack) Pop() Object {
	if s.Len() == 0 {
		return nil
	}
	obj := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return obj
}
