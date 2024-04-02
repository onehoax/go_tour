package myarray

import "fmt"

func MyArrays() {
	var a [2]string
	fmt.Println(a)

	a[0] = "Hello"
	a[1] = "world"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var primes [6]int = [6]int{2, 3, 4, 5, 6}
	fmt.Println(primes)
}
