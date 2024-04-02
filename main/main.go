package main

import (
	"fmt"
	"math"
	"onehoax/go_tour/basic/flowcontrol"
	"onehoax/go_tour/basic/function"
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

	f := function.Fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
