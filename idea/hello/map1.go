package main

import (
	"fmt"
	"golang.org/x/tour/wc"
)

type V struct {
	Lat, Long float64
}

var m map[string]V
var m1 = map[string]V{
	"b l": V{4.1, -2.3},
	"g":   V{2.2, -9},
}

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	m = make(map[string]V)
	m["b l"] = V{4.1, -2.3}
	fmt.Println(m["b l"])
	fmt.Println(m1)

	m2 := make(map[string]int)
	m2["A"] = 42
	fmt.Println(m2["A"])
	m2["A"] = 48
	fmt.Println(m2["A"])

	delete(m2, "A")
	fmt.Println(m2["A"])

	v, ok := m2["A"]
	fmt.Println(v, ok)

	wc.Test(WordCount)
}
