package main

import "fmt"

func main() {
	// var i8 int8
	// fmt.Printf("int8: %T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)

	// var sdh string = `
	// a
	// b
	// c
	// `
	// fmt.Printf("sdh: %v\n", sdh)

	// x := [...]int{1, 2, 3}
	// for _, v := range x {
	// 	fmt.Printf("v: %v\n", v)
	// }
	m := make(map[string]string, 0)
	m["ldd"] = "me"
	for k, v := range m {
		fmt.Printf("k: %v v: %v\n", k, v)
	}

}
