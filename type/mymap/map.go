package mymap

import (
	"fmt"
	"strings"
)

type vertex struct {
	Lat, Long float64
}

var m1 map[string]vertex

func MyMap() {
	m1 = make(map[string]vertex)

	m1["Bell Labs"] = vertex{40.68, -74.76}

	fmt.Println(m1)
	fmt.Println(m1["Bell Labs"])

	// map literal
	m2 := map[string]vertex{
		"Bell Labs": {40.68, -74.76},
		"Google":    {37.98, -122.65},
	}

	fmt.Println(m2)
}

func MapManipulation() {
	m := map[string]int{
		"monday":    1,
		"tuesday":   2,
		"wednesday": 3,
		"thursday":  4,
		"friday":    5,
		"saturday":  6,
	}

	m["sunday"] = 7

	fmt.Println(m)

	elem1, ok := m["monday"]
	fmt.Println("The value:", elem1, "Present?", ok)

	delete(m, "monday")

	elem2, ok := m["monday"]
	fmt.Println("The value:", elem2, "Present?", ok)

	fmt.Println(m)

}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)

	for _, v := range words {
		e, ok := m[v]

		if ok {
			m[v] = e + 1
		} else {
			m[v] = 1
		}
	}

	return m
}
