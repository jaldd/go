package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {
	go func() {
		chanInt <- 100
		chanStr <- "hello"
		close(chanInt)
		close(chanStr)
	}()
	for {
		select {
		case r := <-chanInt:
			fmt.Println("chanInt %v", r)
		case r := <-chanStr:
			fmt.Println("chanStr %v", r)
		default:
			fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
}
