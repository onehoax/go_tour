package function

import "math"

func Swap(x, y string) (string, string) {
	return y, x
}

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

var Hypot = func(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// closure
func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fib closure
func Fib() func() int {
	p1, p2 := 0, 1

	return func() int {
		tmp := p1
		p1 = p2
		p2 = tmp + p2
		return p2
	}
}
