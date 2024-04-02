package main

import (
	"fmt"
	"math"
	"onehoax/go_tour/basic/flowcontrol"
	"onehoax/go_tour/basic/function"
	"onehoax/go_tour/generic/mygeneric"
	"onehoax/go_tour/methodinterface/myinterface"
	"onehoax/go_tour/methodinterface/mymethod"
	"onehoax/go_tour/type/myarray"
	"onehoax/go_tour/type/mymap"
	"onehoax/go_tour/type/mypointer"
	"onehoax/go_tour/type/myslice"
	"onehoax/go_tour/type/mystruct"

	"golang.org/x/tour/wc"
)

func main() {
	fmt.Println("################# functions #################")
	a, b := function.Swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(function.Split(10))

	fmt.Println()

	fmt.Println("################# flow control #################")
	flowcontrol.ForLoop(5)
	flowcontrol.WhileLoop(5)

	fmt.Println()

	fmt.Println(flowcontrol.CustomSqrt(9))
	fmt.Println(
		flowcontrol.CustomPow(3, 2, 10),
		flowcontrol.CustomPow(3, 3, 20),
	)

	fmt.Println()

	flowcontrol.SwitchConstrcut()
	flowcontrol.DaysOfWeek()
	flowcontrol.SwitchNoCond()

	fmt.Println()

	flowcontrol.MyDefer()
	flowcontrol.DeferStack()

	fmt.Println()

	fmt.Println("################# types #################")
	mypointer.MyPointers()

	fmt.Println()

	var v mystruct.Vertex = mystruct.Vertex{X: 1, Y: 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(v.Y)

	var p *mystruct.Vertex = &v
	p.X = 50
	(*p).Y = 79
	fmt.Println(v)
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&p)

	fmt.Println()

	var (
		v1 = mystruct.Vertex{X: 1, Y: 2} // has type Vertex
		v2 = mystruct.Vertex{X: 1}       // Y: 0 is implicit
		v3 = mystruct.Vertex{}           // X: 0 and Y: 0
		pt = &mystruct.Vertex{X: 5, Y: 7}
	)

	fmt.Println(v1, v2, v3, pt)
	fmt.Println()
	myarray.MyArrays()
	fmt.Println()
	myslice.MySlices()
	fmt.Println()
	myslice.MyOtherSlices()
	myslice.SliceMake()
	fmt.Println()
	myslice.Slice2D()
	fmt.Println()
	myslice.SliceAppend()
	fmt.Println()
	myslice.MyRange()

	fmt.Println()

	mymap.MyMap()
	fmt.Println()

	mymap.MapManipulation()

	fmt.Println()
	fmt.Println(mymap.WordCount("hello world hi hey world hi hello hello hey hey hey"))
	wc.Test(mymap.WordCount)

	fmt.Println()
	fmt.Println(function.Hypot(3, 4))
	fmt.Println(function.Compute(function.Hypot))
	fmt.Println(function.Compute(math.Pow))

	fmt.Println()

	pos, neg := function.Adder(), function.Adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(1), neg(-2*i))
	}

	fmt.Println()

	var f func() int = function.Fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	fmt.Println()

	fmt.Println("################# methods & interfaces #################")
	vertex := mymethod.Vertex{X: 3, Y: 4}
	fmt.Println(vertex.AbsReceiver())
	fmt.Println(mymethod.AbsNoReceiver(vertex))
	var myfloat mymethod.MyFloat = mymethod.MyFloat(-math.Sqrt2)
	fmt.Println(myfloat.AbsMyFloat())

	fmt.Println("Method with value reciver works on a copy of arg")
	fmt.Println(vertex)
	vertex.ScaleValue(5.0)
	fmt.Println(vertex)

	fmt.Println("Method with pointer reciver works on the original arg object")
	fmt.Println(vertex)
	vertex.ScalePointer(5.0) // Go interprets as `(&vertex).ScalePointer(5.0)` for convinience
	fmt.Println(vertex)
	(&vertex).ScalePointer(5.0)
	fmt.Println(vertex)

	mymethod.ScaleFunc(&vertex, 2.0)
	fmt.Println(vertex)

	fmt.Println()

	fmt.Println("INTERFACES")

	myinterface.InterfaceImplement()

	fmt.Println()
	fmt.Println("DYNAMIC TYPES")
	myinterface.Test()

	fmt.Println()
	fmt.Println("EMPTY INTERFACES")
	myinterface.EmptyInterface()

	fmt.Println()
	fmt.Println("TYPE ASSERTION")
	myinterface.Assert()

	fmt.Println()
	fmt.Println("TYPE SWITCH")
	myinterface.TestType()

	fmt.Println()
	fmt.Println("STRINGER")
	myinterface.PrintPpl()
	myinterface.PrintIP()

	fmt.Println()
	fmt.Println("ERROR")
	if err := myinterface.RunError(); err != nil {
		fmt.Println(err)
	}

	r, e := myinterface.Divide(10, 2)
	fmt.Println(r, e)
	r, e = myinterface.Divide(10, 0)
	fmt.Println(r, e)

	fmt.Println()
	fmt.Println("READER")
	myinterface.Reader()
	myinterface.Readcustom()
	fmt.Println()

	fmt.Println("################# generics #################")
	mygeneric.Compare()
}
