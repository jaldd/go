package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("A: %v\n", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("B: %v\n", i)
	}
}
func main() {

	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(5)
	go a()
	go b()
	time.Sleep(time.Second)
}
