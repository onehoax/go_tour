package main

import (
	"fmt"
	"onehoax/go_tour/basic/function"
)

func main() {
	a, b := function.Swap("hello", "world")
	fmt.Println(a, b)
}
