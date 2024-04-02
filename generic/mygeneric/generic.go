package mygeneric

import "fmt"

// generic type
type List[T any] struct {
	Next *List[T]
	Val  T
}

// generic function
func index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func Compare() {
	si := []int{10, 20, 15, -10}
	fmt.Println(index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(index(ss, "hello"))
}
