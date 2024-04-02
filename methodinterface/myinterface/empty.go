package myinterface

import "fmt"

// interface that implements 0 methods is the empty interface
// an empty interface may hold values of any type (every type implements >= 0 methods)
// empty interfaces are used by code that handles values of unknown type -> e.g.: `fmt.Print` takes any number of args of type `interface{}`
func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func EmptyInterface() {
	var i interface{}
	describeEmpty(i)

	i = 42
	describeEmpty(i)

	i = "Hello"
	describeEmpty(i)
}
