package main

import (
	"fmt"
	"runtime"
	"time"
)

func show(s string) {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			runtime.Goexit()
		}
		fmt.Printf("i: %v\n", i)
	}
}
func main() {

	go show("hello")
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		time.Sleep(time.Second)
	}
}
