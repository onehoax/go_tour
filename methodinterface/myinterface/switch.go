package myinterface

import "fmt"

func switchType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("value %v, type %T\n", v, v)
	case string:
		fmt.Printf("value %v, type %T\n", v, v)
	default:
		fmt.Printf("I Don't know about type %T\n", v)
	}
}

func TestType() {
	switchType(21)
	switchType("Hello")
	switchType(true)
}
