package mymethod

import (
	"math"
)

// you can only declare a mthod with a receiver whose type is defined in the same package as the method
type MyFloat float64

func (f MyFloat) AbsMyFloat() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) AbsReceiver() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsNoReceiver(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// value receiver; operates on a copy of the original Vertex object
// called as `v.ScaleValue(x)`
// can also call as `(&v).ScaleValue(x)` but the method will still only work on a copy of the arg
func (v Vertex) ScaleValue(f float64) {
	v.X *= f
	v.Y *= f
}

// pointer receiver: operates on the original Vertex object
// correct way of calling is `(&v).ScalePointer(x)` but can call as `v.ScalePointer(x)` as go translates it to `(&v).ScalePointer(x)` for convienience
func (v *Vertex) ScalePointer(f float64) {
	v.X *= f
	v.Y *= f
}

// needs to be called as `ScalePointer(&v, x)`
func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}
