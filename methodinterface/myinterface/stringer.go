package myinterface

import "fmt"

type person struct {
	name string
	age  int
}

// implement interface `Stringer` from `fmt` to override how values get printed for type *person
func (p *person) String() string {
	return fmt.Sprintf("name: %v, age: %v", p.name, p.age)
}

func PrintPpl() {
	a := &person{name: "andres osorio", age: 33}
	b := &person{name: "valen osorio", age: 26}
	fmt.Println(a)
	fmt.Println(b)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func PrintIP() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googledns": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
