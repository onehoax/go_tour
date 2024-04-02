package mypointer

import "fmt"

func MyPointers() {
	i, j := 42, 2701

	var p *int = &i // p is of type pointer to int
	fmt.Printf("reading i through itself as value: %d\n", i)
	fmt.Printf("reading i through dereference of p: %d\n", *p)
	fmt.Printf("reading address of i through p: %p\n", p)
	fmt.Printf("reading address of i through i: %p\n", &i)
	fmt.Printf("reading address of p: %p\n", &p)

	*p = 21 // set i through the pointer
	fmt.Println(i)
	i = 8 // set i through itself
	fmt.Println(*p)

	p = &j // point to j
	fmt.Println(j)
	fmt.Println(*p)
	*p /= 2

	fmt.Println(j)
	fmt.Println(*p)
}
