package myslice

import (
	"fmt"
	"strings"
)

func MySlices() {
	primes := []int{2, 3, 5, 7, 11, 13}

	// low bound is inclusive, high bound is exclusive
	var slice []int = primes[1:4]

	fmt.Println(primes)
	fmt.Println(slice)

	names := []string{
		"jhon",
		"paul",
		"george",
		"ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	z := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(z)

	t := z[1:4]
	fmt.Println(t)

	// default values for low/high bounds are 0 and slice's length respectively
	t = z[:4]
	fmt.Println(t)

	t = z[3:]
	fmt.Println(t)

	t = z[:]
	fmt.Println(t)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	if s == nil {
		fmt.Println("nil!")
	}
}

func MyOtherSlices() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// give it 0 length
	s = s[:0]
	printSlice(s)

	// extend its length
	s = s[:4]
	printSlice(s)

	// drop 1st 2 values
	s = s[2:]
	printSlice(s)

	s = s[:cap(s)]
	printSlice(s)

	var q []int
	printSlice(q)

}

func SliceMake() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 5, 10)
	printSlice(b)
}

func Slice2D() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func SliceAppend() {
	// new just allocates memory; make allocates and initializes memmory
	s := *new([]int) // same as `var s []int`
	printSlice(s)

	// append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// the slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// we can add more than 1 element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)

	a := []string{"jhon", "paul"}
	b := []string{"george", "ringo", "pete"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	fmt.Println("APPEND", a)
}

func MyRange() {
	a := []int{1, 2, 3, 4}

	for i, v := range a {
		fmt.Printf("index=%d value=%d\n", i, v)
	}

	// only want index
	for i := range a {
		fmt.Printf("index=%d\n", i)
	}

	// only want value
	for _, v := range a {
		fmt.Printf("value=%d\n", v)
	}

	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << i // 2**i
	}

	for i, v := range pow {
		fmt.Printf("index=%d value=%d\n", i, v)
	}

}
