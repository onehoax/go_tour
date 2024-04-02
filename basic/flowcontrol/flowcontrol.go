package flowcontrol

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// regular for loop
func ForLoop(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("for loop: %d\n", i)
	}
}

// while loop
func WhileLoop(n int) {
	i := 0
	for i < n {
		i++
		fmt.Printf("while loop: %d\n", i)
	}
}

func CustomSqrt(x float64) string {
	if x < 0 {
		return CustomSqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func CustomPow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func SwitchConstrcut() {
	fmt.Print("Go is running on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n", os)
	}
}

func DaysOfWeek() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	saturday := time.Saturday
	fmt.Printf("today is %s (%d) and type is %T\n", today, today, today)
	fmt.Printf("saturday is %s (%d) and type is %T\n", saturday, saturday, saturday)

	switch saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

// switch without a condition is the same as `switch true`
// clean way to write long `if-then-else` chains
func SwitchNoCond() {
	t := time.Now()

	fmt.Println(t.Hour())

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

// a defer statement defers the execution of a function until the surrounding function returns
// the deferred call's args are evaluated immediately, but the function call is not executed until the surrounding function returns
func MyDefer() {
	val := "initial value"
	defer fmt.Printf("finishing execution; val: %s\n", val)

	val = "modified value"
	fmt.Printf("starting execution; val: %s\n", val)
}

// deferred function calss are pushed onto a stack
// when a function returns, its deferred calls are executed in a LIFO
func DeferStack() {
	fmt.Println("counting")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
