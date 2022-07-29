package main

import "fmt"

type V struct {
	Lat, Long float64
}

var m map[string]V
var m1 = map[string]V{
	"b l": V{4.1, -2.3},
	"g":   V{2.2, -9},
}

func main() {
	m = make(map[string]V)
	m["b l"] = V{4.1, -2.3}
	fmt.Println(m["b l"])
	fmt.Println(m1)

}
