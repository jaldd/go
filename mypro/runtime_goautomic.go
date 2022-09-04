package main

import (
	"fmt"
	"sync/atomic"
)

var j int32 = 100

func add() {
	atomic.AddInt32(&j, 1)
}
func sub() {
	atomic.AddInt32(&j, -1)
}
func main() {

	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	fmt.Printf("%v\n", j)
}
