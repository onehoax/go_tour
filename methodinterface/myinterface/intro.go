package myinterface

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func InterfaceImplement() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{X: 3, Y: 4}

	a = f // MyFloat implements Abser -> ok
	fmt.Println(a.Abs())

	a = &v // *Vertex implements Abser -> ok
	// a = v  Error bc Vertex does not implement Abser, only *Vertex does

	fmt.Println(a.Abs())
}
