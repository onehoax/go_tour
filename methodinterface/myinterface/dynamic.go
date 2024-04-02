package myinterface

import (
	"fmt"
	"math"
)

type I interface {
	m()
}

type t struct {
	s string
}

func (t *t) m() {
	if t == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(t.s)
}

type f float64

func (f f) m() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func Test() {
	var i I

	i = &t{"Hello"}
	describe(i)
	i.m()

	i = f(math.Pi)
	describe(i)
	i.m()

	var q *t
	i = q
	describe(i)
	i.m()
}
