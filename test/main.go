package main

import (
	"fmt"
	"log"
	"reflect"
)

var g = initGrim()

// Grim state machine
type grim struct {
	callStack []*frame
	stack     []*any
	heap      []*variable
}

type frame struct {
	name       string
	stackStart uint
	heapStart  uint
}

type variable struct {
	name  string
	tp    reflect.Kind
	value any
}

func initGrim() grim {
	return grim{}
}

func (g *grim) trace() {
	fmt.Print("stack:\n")
	for i := 0; i < len(g.stack); i++ {
		fmt.Printf("[%d] = %#v\n", i, g.stack[i])
	}
	fmt.Print("heap:\n")
	for i := 0; i < len(g.heap); i++ {
		fmt.Printf("[%d] = %#v\n", i, g.heap[i])
	}
	fmt.Print("call stack:\n")
	for i := 0; i < len(g.callStack); i++ {
		fmt.Printf("[%d] = %#v\n", i, g.callStack[i])
	}
}

func (g *grim) saveVariable(v *variable) {
	g.heap = append(g.heap, v)
}

func var_new(name string, tp reflect.Kind) {
	v := &variable{name, tp, nil}
	g.saveVariable(v)
}

func var_push(name string) {
	v := g.getVariable(name)
	if v == nil {
		log.Fatalf("variable %s not found\n", name)
	}
}

// ----------------------------------------------------------------

func main() {
	g.var_new("n", reflect.Int)
	g.trace()
}
