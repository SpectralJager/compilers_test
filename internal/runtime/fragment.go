package runtime

type Frame struct {
	Start    int
	End      int
	Name     string
	Varibles []Object
}
